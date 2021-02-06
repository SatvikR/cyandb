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
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var queryRunner *QueryRunner

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// homePage: route for home page
func homePage(w http.ResponseWriter, _ *http.Request) {
	if _, err := fmt.Fprintf(w, "Welcome to CyanDB"); err != nil {
		log.Fatal(err)
	}
}

// reader handles the websocket messages
func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("COMMAND RECEIVED: " + string(p))

		if err := conn.WriteMessage(websocket.TextMessage, queryRunner.RunQuery(Parse(p))); err != nil {
			log.Println(err)
			return
		}
	}
}

// wsEndpoint handles websocket connections
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Successfully Connected")

	reader(ws)
}

// setupRoutes sets up routes
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

// StartServer starts the websocket server
func (server *Server) StartServer() {
	queryRunner = NewQueryRunner(server)

	fmt.Println("WEBSOCKETS")
	setupRoutes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.Port), nil))
}
