package main

import (
	"LoginModuleServer/init"
	"LoginModuleServer/login"
	"fmt"
)

func main() {
	err := init.Initialize()
	if err != nil {
		fmt.Println("init err")
		return
	}

	login.Start()
}
