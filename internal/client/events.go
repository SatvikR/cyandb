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

package client

import (
	"github.com/sacOO7/gowebsocket"
	"log"
	"os"
)

var messagePrinter *MessagePrinter

func OnDisconnected(_ error, _ gowebsocket.Socket) {
	log.Println("Disconnected from server")
	os.Exit(1)
}

func OnConnected(_ gowebsocket.Socket) {
	log.Println("Connected to server")
}

func OnConnectError(err error, _ gowebsocket.Socket) {
	log.Println("Received error: ", err)
	os.Exit(1)
}

func OnTextMessage(message string, _ gowebsocket.Socket) {
	messagePrinter.Messages <- message
}
