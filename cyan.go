package main

import (
	"fmt"
	"github.com/SatvikR/cyan-db/server"
	"log"
)

func main() {
	key, err := server.Get("password")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(key)

/*	fmt.Println(server.Set("username", "kevin"))
	fmt.Println(server.Set("password", "kevins_password"))
	fmt.Println(server.Set("likes", "100"))
*/}
