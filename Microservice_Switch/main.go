package main

import (
	"Microservice_Switch/Initialize"
	"Microservice_Switch/server"
	"fmt"
)

func main() {
	err := Initialize.Initialize()
	if err != nil {
		fmt.Println("Initialize err", err)
		return
	}

	server.ServerRouter()
}
