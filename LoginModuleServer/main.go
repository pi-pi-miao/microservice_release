package main

import (
	"LoginModuleServer/login"
	"fmt"
	//"LoginModuleServer/init"
	"LoginModuleServer/init"
)

func main() {
	err := init.Initialize()
	if err != nil {
		fmt.Println("init err")
		return
	}

	login.Start()
}
