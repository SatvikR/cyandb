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

// CyanDB is a simple key value database
//
// Server:
//
// Import server with the "github.com/SatvikR/cyandb/server" package
//
// The keys and values must have 32 bit lengths. There are plans to make keys a
// maximum of 16 bits, because no one practically uses very large keys
//
// The server struct is a container for all of the server specific data
// that includes the storage location and websocket info
//
// Create a server with the CreateServer method, ex.
//
//	db := server.CreateServer(server.DefaultDBPath, server.DefaultPort)
//
// server.DefaultDBPath and server.DefaultPort can be changed with a custom path if needed
//
// As of right now, there are two commands: Get and Set
//
// Set creates a key value pair. The data is serialized like so:
// 4 bytes for len(key), 4 bytes for len(val), len(key) bytes for key, len(val) bytes for val.
// Right now, if a duplicate key exists, it creates the new pair anyway, and ergo, the new pair will be ignored.
// Set returns the value
// ex.
//
//	output, err := db.Set("hello", "world")
//
// Get gets a value given a key. Get returns an empty string and an io.EOF error if the key doesn't exist. In the future,
// Keys will be stored in an index file, and values will be stored separately, so this method will be significantly
// faster.
// ex.
//
// 	output, err, _ := db.Get("hello")
//
// More practically, the commands will be called via a websocket client. The websockets are handled in ./server/websocket.go.
// The websocket messages are passed into a json parser, which return a Query struct. The queries are passed into
// a query runner which runs the command.
// Client:
//
//
// A simple terminal based websocket client. Still in progress
package main
