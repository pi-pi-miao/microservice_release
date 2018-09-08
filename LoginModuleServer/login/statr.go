package login

import (
	"LoginModuleServer/controller"
	"net"
	"net/rpc"
)

func Start() {
	rect := new(controller.U)
	rpc.Register(rect)
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	controller.ChkError(err)
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	controller.ChkError(err2)

	for {
		conn, err3 := tcplisten.Accept()
		if err3 != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
