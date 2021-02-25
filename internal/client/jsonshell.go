//
//    Copyright 2021 Satvik Reddy
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

package client

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	singleQuote = iota
	doubleQuote = iota
	noQuote     = iota
)

type Stack struct {
	Data []rune
}

// pop pops from the stack
func (stack *Stack) pop() {
	stack.Data = stack.Data[:len(stack.Data)-1]
}

// push pushes to the stack
func (stack *Stack) push(r rune) {
	stack.Data = append(stack.Data, r)
}

// top returns the top element of the stack
func (stack *Stack) top() rune {
	return stack.Data[len(stack.Data)-1]
}

// GetJsonInput gets the json command from the user
func GetJsonInput(promptOne string, promptTwo string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	jsonStack := &Stack{Data: make([]rune, 0)}

	inString := false
	quoteType := noQuote

	currentInput := ""

	tokenMap := make(map[rune]rune)
	tokenMap[')'] = '('
	tokenMap['}'] = '{'
	tokenMap[']'] = '['

	prompt := promptOne

	for {
		fmt.Printf(prompt)
		line, _ := reader.ReadString('\n')

		if len(line) > 4 {
			if line[:4] == "exit" {
				os.Exit(0)
			}
		}

		currentInput += line

		for _, char := range line {
			if char == '"' || char == '\'' {
				if char == '"' {
					if inString && (quoteType == doubleQuote) {
						inString = !inString
					} else if !inString {
						inString = !inString
						quoteType = doubleQuote
					}
				} else {
					if inString && (quoteType == singleQuote) {
						inString = !inString
					} else if !inString {
						inString = !inString
						quoteType = singleQuote
					}
				}
			} else if !inString {
				if char == '(' || char == '{' || char == '[' {
					jsonStack.push(char)
				} else if char == ')' || char == '}' || char == ']' {
					if jsonStack.top() != tokenMap[char] {
						return "", errors.New("invalid JSON")
					} else {
						jsonStack.pop()
					}
				}
			}
		}

		if len(jsonStack.Data) == 0 {
			return currentInput, nil
		} else {
			prompt = "..."
		}
	}
}
