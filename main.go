package main

import (
	//"github/Enocp/webapi-go/server"

	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}