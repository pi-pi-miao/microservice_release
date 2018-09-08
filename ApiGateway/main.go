package main

import (
	"fmt"
	"net/http"
	"net"
	"golang.org/x/net/netutil"
	"time"
	"ApiGateway/Proxy"
)


func main() {
	Mux:=http.NewServeMux()
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err listen")
	}
	defer l.Close()
	L := netutil.LimitListener(l, 10000)


	Mux.HandleFunc("/login",Proxy.Login)

	//待定
	Mux.HandleFunc("/api",Proxy.SelectApi)

	Mux.HandleFunc("/signup",Proxy.Register)

	Mux.HandleFunc("/signup_email",Proxy.RegisterEmail)


	server:=&http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:Mux,
	}
	server.Serve(L)

}
