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

package main

import (
	"fmt"
	"github.com/SatvikR/cyandb/internal/server"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		db := server.CreateServer(server.DefaultDBPath, server.DefaultPort)

		db.StartServer()
		os.Exit(1)
	} else {
		if len(args) != 1 {
			fmt.Println("Invalid arguments. Try `cyand help`")
			os.Exit(1)
		} else if args[0] == "help" {
			fmt.Println("Usage:\n" +
				"	cyand [options]\n" +
				"Info: \n" +
				"	Running `cyand` on its own will start a cyandb server at ws://localhost:8080/ws\n" +
				"Options: \n" +
				"	help	Prints this message")
		} else {
			fmt.Println("Invalid arguments. Try `cyand help`")
			os.Exit(1)
		}
	}
}
