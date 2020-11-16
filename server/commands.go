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
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Set adds a key value pair to a database
// Serialization schema:
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

	f, _ := os.OpenFile(server.Location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

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
func Get(key string) (string, error) {
	byteFileContents, _ := ioutil.ReadFile(DefaultDBPath)
	// Convert file into string
	fileContents := string(byteFileContents)

	// i is the current index in the file
	i := 0

	for {
		curr := ""

		// Get current key/value string into curr
		currentSlice := fileContents[i:]
		for charIndex := range currentSlice {
			char := currentSlice[charIndex]
			// Break here because ';' is our record separator
			if char == ';' {
				break
			}
			curr += string(char)
		}

		// Get current key
		keyLen, _ := strconv.ParseInt(string(curr[0]), 10, 64)
		currKey := curr[1 : keyLen+1]
		// Delete key from curr
		curr = curr[keyLen+1:]
		// Get current value
		valueLen, _ := strconv.ParseInt(string(curr[0]), 10, 64)
		currVal := curr[1 : valueLen+1]
		// Return if key is found
		if key == currKey {
			return currVal, nil
		}
		// Move i along to next pair
		i += int(keyLen) + int(valueLen) + 3 // 3 represents three characters :, ;, and the first char in the next pair
		if i >= len(fileContents) {
			return "", errors.New("key not found")
		}
	}
}
