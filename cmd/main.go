package main

import (
	"log"

	"github.com/ChiragRajput101/rest-api/cmd/api"
)

// instantiate the server
func main() {
	server := api.InitServer(":8000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}