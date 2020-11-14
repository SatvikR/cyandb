package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	DBFile = "/data/cyan.db"
)

func Set(key string, val string) string {
	rawFileContents, _ := ioutil.ReadFile(DBFile)
	out := []byte(fmt.Sprintf("%s%d%s%d%s;", string(rawFileContents),len(key), key, len(val), val))

	err := ioutil.WriteFile(DBFile, out, 0664)

	if err != nil {
		log.Fatal(err)
	}

	return val
}

// Get returns a value from the database given a corresponding key
func Get(key string) (string, error) {
	byteFileContents, _ := ioutil.ReadFile(DBFile)
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
		keyLen, _ := strconv.ParseInt(string(curr[0]),10, 64)
		currKey := curr[1:keyLen + 1]
		// Delete key from curr
		curr = curr[keyLen + 1:]
		// Get current value
		valueLen, _ := strconv.ParseInt(string(curr[0]), 10, 64)
		currVal := curr[1:valueLen + 1]
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
