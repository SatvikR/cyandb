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
	"encoding/json"
	"fmt"
)

type Query struct {
	Command string
	Args    []string
}

type QueryRunner struct {
	Server *Server
}

func NewQueryRunner(server *Server) *QueryRunner {
	return &QueryRunner{Server: server}
}

// NewQuery generates a query
func NewQuery(command string, args []string) *Query {
	return &Query{Command: command, Args: args}
}

// Parse parses a command
func Parse(text []byte) (error, *Query) {
	var query Query
	if err := json.Unmarshal(text, &query); err != nil {
		return err, nil
	}

	return nil, NewQuery(query.Command, query.Args)
}

func VerifyGetQuery(query *Query) ([]byte, bool) {
	if len(query.Args) != 1 {
		return []byte("Error: Invalid arguments"), true
	}
	return nil, false
}

func VerifySetQuery(query *Query) ([]byte, bool) {
	if len(query.Args) != 2 {
		return []byte("Error: Invalid arguments"), true
	}
	return nil, false
}

func (queryRunner *QueryRunner) RunQuery(err error, query *Query) []byte {
	if err != nil {
		return []byte("Error: Invalid command")
	}

	if query.Command == "set" {
		bytes, done := VerifySetQuery(query)
		if done {
			return bytes
		}

		out, err := queryRunner.Server.Set(query.Args[0], query.Args[1])
		if err != nil {
			return []byte(fmt.Sprintf("SET %s %s; OUTPUT=%s", query.Args[0], query.Args[1], err))
		} else {
			return []byte(fmt.Sprintf("SET %s %s; OUTPUT=%s", query.Args[0], query.Args[1], out))
		}
	} else if query.Command == "get" {
		bytes, done := VerifyGetQuery(query)
		if done {
			return bytes
		}
		out, err, _ := queryRunner.Server.Get(query.Args[0])
		if err != nil {
			return []byte(fmt.Sprintf("GET %s; OUTPUT=%s", query.Args[0], err))
		} else {
			return []byte(fmt.Sprintf("GET %s; OUTPUT=%s", query.Args[0], out))
		}
	} else {
		return []byte("Invalid command")
	}
}
