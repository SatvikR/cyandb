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

// Internal code for the cyandb server
package server

import (
	"os"
)

const (
	// Use this for default db functionality
	DefaultDBPath = "/data/cyan/"
	DefaultPort   = 8080
	dataFileName  = "db.dat"
)

// Server is the struct definition for the server
// I will add more to this once websockets are introduced
type Server struct {
	Location string
	Port     int
}

// CreateServer creates a server struct
func CreateServer(path string, port int) *Server {
	// Make sure path ends in '/'
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	// Create db directory if doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		// log.Fatal(err)
	}

	server := &Server{Location: path + dataFileName, Port: port}

	return server
}
