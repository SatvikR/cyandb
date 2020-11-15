package main

import (
	"fmt"
	"github.com/SatvikR/cyandb/server"
	"log"
)

func main() {
	key, err := server.Get("password")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(key)
}
