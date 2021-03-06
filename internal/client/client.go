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

// Internal code for the cyandb client
package client

import (
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"os"
	"os/signal"
)

const (
	DefaultAddr = "ws://localhost:8080/ws"
)

type Client struct {
	Addr   string
	Socket gowebsocket.Socket
}

type MessagePrinter struct {
	Messages chan string
}

// CreateClient creates a client
func CreateClient(addr string) *Client {
	client := &Client{Addr: addr, Socket: gowebsocket.New(addr)}
	client.setupEvents()
	return client
}

// StartClient will be updated later
func (client *Client) StartClient() {
	messagePrinter = &MessagePrinter{Messages: make(chan string)}
	fmt.Println("Starting client..., type `exit` to exit")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	client.Socket.Connect()

	for {
		line, err := GetJsonInput(fmt.Sprintf("%s>", DefaultAddr), "...")
		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			client.Socket.SendText(line[:len(line)-1])

			res := <-messagePrinter.Messages
			fmt.Println(res)
		}

	}
}

// setupEvents ties callback methods to the event handlers
func (client *Client) setupEvents() {
	client.Socket.OnConnected = OnConnected
	client.Socket.OnDisconnected = OnDisconnected
	client.Socket.OnConnectError = OnConnectError
	client.Socket.OnTextMessage = OnTextMessage
}
