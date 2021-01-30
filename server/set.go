//
//  Copyright 2021 Satvik Reddy
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
	"io"
	"os"
)

// SetExisting Sets value for an existing key
func (server *Server) SetExisting(keyAsBytes []byte, valAsBytes []byte, lenKey uint32, lenVal uint32,
	originalValue string, existingPos int64) error {

	f, err := os.OpenFile(server.Location, os.O_APPEND|os.O_RDWR, filePermissions)

	if err != nil {
		return err
	}

	endBuffer, err := getEndBuffer(f, originalValue, existingPos)
	if err != nil {
		_ = f.Close()
		return err
	}

	_, err = f.Seek(existingPos, io.SeekStart)
	if err != nil {
		_ = f.Close()
		return err
	}

	err = replaceEndBuffer(f, existingPos, keyAsBytes, valAsBytes, endBuffer, lenKey, lenVal)
	if err != nil {
		_ = f.Close()
		return err
	}

	if err = f.Close(); err != nil {
		return err
	}

	return err
}

// getEndBuffer gets bytes from end of existing value to the EOF
func getEndBuffer(f *os.File, originalValue string, existingPos int64) ([]byte, error) {
	// Seek to end of original key
	_, err := f.Seek(existingPos, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Seek to end of original value
	endOfOriginalVal, err := f.Seek(int64(len(originalValue)), io.SeekCurrent)
	if err != nil {
		return nil, err
	}

	// Find EOF value (len of file)
	eof, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	// Seek back to where we were before
	_, err = f.Seek(endOfOriginalVal, io.SeekStart)
	if err != nil {
		return nil, err
	}

	endBuffer := make([]byte, eof-endOfOriginalVal)

	// Read rest of file into a buffer
	_, err = f.Read(endBuffer)
	if err != nil {
		return nil, err
	}

	return endBuffer, err
}

// replaceEndBuffer replaces the endBuffer with the new key and value
func replaceEndBuffer(f *os.File, existingPos int64, keyAsBytes []byte, valAsBytes []byte, endBuffer []byte, lenKey uint32, lenVal uint32) error {
	// Go to 4 bytes behind the original key (4 bytes is the length in bytes of the length of the keys/values,
	// see line 32
	err := f.Truncate(existingPos - int64(lenKey) - 4)
	if err != nil {
		return err
	}

	_, err = f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	// Write new length of value and original key
	_, err = f.Write(append(Uint32ToByteArr(lenVal), keyAsBytes...))
	if err != nil {
		return err
	}

	// Write the new value and the buffer of the rest of the file
	_, err = f.Write(valAsBytes)
	if err != nil {
		return err
	}
	_, err = f.Write(endBuffer)
	if err != nil {
		return err
	}

	return err
}
