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
	"github.com/SatvikR/cyandb/internal/client"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		dbClient := client.CreateClient(client.DefaultAddr)
		dbClient.StartClient()
		os.Exit(1)
	} else {
		if len(args) != 1 {
			fmt.Println("Invalid arguments. Try `cyansh help`")
			os.Exit(1)
		} else if args[0] == "help" {
			fmt.Println("Usage:\n" +
				"	cyansh [options]\n" +
				"Info: \n" +
				"	Running `cyansh` on its own will cyandb shell\n" +
				"Options: \n" +
				"	help	Prints this message")
		} else {
			fmt.Println("Invalid arguments. Try `cyansh help`")
			os.Exit(1)
		}
	}
}
