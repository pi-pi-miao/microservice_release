package main

import (
	"ApiGateway/initialize"
	"ApiGateway/server"
)

func main() {

	initialize.Initall()

	server.Start()
}
