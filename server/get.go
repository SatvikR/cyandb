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

// GetAsBytes gets the key and value as bytes
func GetAsBytes(f *os.File) ([]byte, []byte, error) {
	// Read 4 bytes to get current key length
	currKeyLenAsBytes := make([]byte, 4)
	_, err := f.Read(currKeyLenAsBytes)
	if err != nil {
		return nil, nil, err
	}

	// Repeat previous step for value
	currValLenAsBytes := make([]byte, 4)
	_, err = f.Read(currValLenAsBytes)
	if err != nil {
		return nil, nil, err
	}

	return currKeyLenAsBytes, currValLenAsBytes, nil
}

// ReadKeyVal reads the key and value from the current point in the file
func ReadKeyVal(currKeyLen uint32, err error, f *os.File, currValLen uint32) (
	error,
	int64,
	string,
	[]byte,
	string,
	error,
	int64,
	bool,
) {
	// Read key using key len from above
	currKeyAsBytes := make([]byte, currKeyLen)
	_, err = f.Read(currKeyAsBytes)
	if err != nil {
		return nil, 0, "", nil, "", err, 0, true
	}

	currentPos, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, 0, "", nil, "", err, 0, true
	}

	// Convert bytes to string
	currKey := string(currKeyAsBytes)

	// Repeat previous step for value
	currValAsBytes := make([]byte, currValLen)
	_, err = f.Read(currValAsBytes)
	if err != nil {
		return nil, 0, "", nil, "", err, 0, true
	}
	return err, currentPos, currKey, currValAsBytes, "", nil, 0, false
}
