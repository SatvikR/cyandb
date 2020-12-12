//
//    Copyright 2020 Satvik Reddy
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package server

import (
	"encoding/binary"
	"io"
	"os"
)

const (
	filePermissions = 0644
)

// Set adds a key value pair to a database
// Serialization model:
//
// [4 bytes len(key)][4 bytes len(val)][len(key) bytes key][len(val) bytes val]
func (server *Server) Set(key string, val string) (string, error) {
	// Get lengths of key and value as unsigned 32 bit integers
	lenKey := uint32(len(key))
	lenVal := uint32(len(val))

	// Get key and value as byte arr
	keyAsBytes := []byte(key)
	valAsBytes := []byte(val)

	oldVal, _, existingPos := server.Get(key)
	if existingPos != 0 {
		tempFile, err := os.OpenFile(server.Location+".temp", os.O_CREATE|os.O_WRONLY, filePermissions)

		if err != nil {
			return "", err
		}
		originalFile, err := os.OpenFile(server.Location, os.O_CREATE|os.O_RDWR, filePermissions)

		if err != nil {
			return "", err
		}

		tempBuffer := make([]byte, existingPos)
		_, err = originalFile.Read(tempBuffer)
		if err != nil {
			return "", err
		}

		tempBuffer = tempBuffer[:0]

		_, err = tempFile.Write(tempBuffer)
		if err != nil {
			return "", err
		}

		currentPos, err := originalFile.Seek(int64(len(oldVal)), io.SeekCurrent)
		if err != nil {
			return "", err
		}

		EOF, err := originalFile.Seek(0, io.SeekEnd)
		if err != nil {
			return "", err
		}
		_, err = originalFile.Seek(currentPos, io.SeekStart)
		if err != nil {
			return "", err
		}

		restOFBuffer := make([]byte, EOF-currentPos)
		_, err = originalFile.Read(restOFBuffer)
		if err != nil {
			return "", err
		}

		newValBuffer := []byte(val)
		_, err = tempFile.Write(newValBuffer)
		if err != nil {
			return "", err
		}

		_, err = tempFile.Write(restOFBuffer)
		if err != nil {
			return "", err
		}

		err = tempFile.Close()
		if err != nil {
			return "", err
		}

		err = originalFile.Close()
		if err != nil {
			return "", err
		}

		err = os.Remove(originalFile.Name())
		if err != nil {
			return "", err
		}
		err = os.Rename(tempFile.Name(), server.Location)
		if err != nil {
			return "", err
		}

		return val, nil
	}

	// Append all bytes into one slice
	// Chaining appends to append more than two slices (This is the only way I could think to do this)
	data := append(Uint32ToByteArr(lenKey), append(Uint32ToByteArr(lenVal), append(keyAsBytes, valAsBytes...)...)...)

	f, _ := os.OpenFile(server.Location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermissions)

	_, err := f.Write(data)
	if err != nil {
		return "", err
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	return val, nil
}

// Get returns a value from the database given a corresponding key
func (server *Server) Get(key string) (string, error, int64) {
	f, _ := os.OpenFile(server.Location, os.O_RDONLY|os.O_CREATE, filePermissions)

	for {
		// Read 4 bytes to get current key length
		currKeyLenAsBytes := make([]byte, 4)
		_, err := f.Read(currKeyLenAsBytes)
		if err != nil {
			return "", err, 0
		}

		// Convert length as bytes into uint32
		currKeyLen := binary.LittleEndian.Uint32(currKeyLenAsBytes)

		// Repeat previous step for value
		currValLenAsBytes := make([]byte, 4)
		_, err = f.Read(currValLenAsBytes)
		if err != nil {
			return "", err, 0
		}

		currValLen := binary.LittleEndian.Uint32(currValLenAsBytes)

		// Read key using key len from above
		currKeyAsBytes := make([]byte, currKeyLen)
		_, err = f.Read(currKeyAsBytes)
		if err != nil {
			return "", err, 0
		}

		currentPos, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			return "", err, 0
		}

		// Convert bytes to string
		currKey := string(currKeyAsBytes)

		// Repeat previous step for value
		currValAsBytes := make([]byte, currValLen)
		_, err = f.Read(currValAsBytes)
		if err != nil {
			return "", err, 0
		}

		// We can continue here if the target key and
		// current key do not match because our file pointer is
		// already at the correct position
		if key != currKey {
			continue
		}

		// Now it is ok to read the value
		currVal := string(currValAsBytes)

		err = f.Close()
		if err != nil {
			return "", err, 0
		}

		return currVal, nil, currentPos
	}
}
