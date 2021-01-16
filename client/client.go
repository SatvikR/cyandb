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
	"bufio"
	"fmt"
	"github.com/sacOO7/gowebsocket"
	"log"
	"os"
	"os/signal"
)

const (
	DefaultAddr = "ws://localhost:8080/ws"
)

type Client struct {
	Addr   string
	socket gowebsocket.Socket
}

// CreateClient creates a client
func CreateClient(addr string) *Client {
	client := &Client{Addr: addr, socket: gowebsocket.New(addr)}
	client.setupEvents()

	return client
}

// StartClient will be updated later
func (client *Client) StartClient() {
	fmt.Println("Starting client...")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	client.socket.Connect()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[%s]> ", DefaultAddr)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if line == "exit" {
			client.socket.Close()
			os.Exit(1)
		}
		client.socket.SendText(line)
	}
}

// setupEvents ties callback methods to the event handlers
func (client *Client) setupEvents() {
	client.socket.OnConnected = OnConnected
	client.socket.OnDisconnected = OnDisconnected
	client.socket.OnConnectError = OnConnectError
}
