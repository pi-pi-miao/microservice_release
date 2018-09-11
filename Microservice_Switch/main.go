package main

import (
	"Microservice_Switch/Initialize"
	"fmt"
	"Microservice_Switch/server"
)

func main(){
	err := Initialize.Initialize()
	if err != nil{
		fmt.Println("Initialize err",err)
		return
	}

	server.ServerRouter()
}

