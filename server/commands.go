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
	"log"
	"os"
)

const (
	filePermissions = 0644
)

// Set adds a key value pair to a database
// Serialization model:
//
// [4 bytes len(key)][4 bytes len(val)][len(key) bytes key][len(val) bytes val]
func (server *Server) Set(key string, val string) string {
	// Get lengths of key and value as unsigned 32 bit integers
	lenKey := uint32(len(key))
	lenVal := uint32(len(val))

	// Get key and value as byte arr
	keyAsBytes := []byte(key)
	valAsBytes := []byte(val)

	// Append all bytes into one slice
	// Chaining appends to append more than two slices (This is the only way I could think to do this)
	data := append(Uint32ToByteArr(lenKey), append(Uint32ToByteArr(lenVal), append(keyAsBytes, valAsBytes...)...)...)

	f, _ := os.OpenFile(server.Location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermissions)

	_, err := f.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	return val
}

// Get returns a value from the database given a corresponding key
func (server *Server) Get(key string) (string, error) {
	f, _ := os.OpenFile(server.Location, os.O_RDONLY|os.O_CREATE, filePermissions)

	for {
		// Read 4 bytes to get current key length
		currKeyLenAsBytes := make([]byte, 4)
		_, err := f.Read(currKeyLenAsBytes)
		if err != nil {
			return "", err
		}

		// Convert length as bytes into uint32
		currKeyLen := binary.LittleEndian.Uint32(currKeyLenAsBytes)

		// Repeat previous step for value
		currValLenAsBytes := make([]byte, 4)
		_, err = f.Read(currValLenAsBytes)
		if err != nil {
			return "", err
		}

		currValLen := binary.LittleEndian.Uint32(currValLenAsBytes)

		// Read key using key len from above
		currKeyAsBytes := make([]byte, currKeyLen)
		_, err = f.Read(currKeyAsBytes)
		if err != nil {
			return "", err
		}

		// Convert bytes to string
		currKey := string(currKeyAsBytes)

		// Repeat previous step for value
		currValAsBytes := make([]byte, currValLen)
		_, err = f.Read(currValAsBytes)
		if err != nil {
			return "", err
		}

		// We can continue here if the target key and
		// current key do not match because our file pointer is
		// already at the correct position
		if key != currKey {
			continue
		}

		// Now it is ok to read the value
		currVal := string(currValAsBytes)

		return currVal, nil
	}
}
