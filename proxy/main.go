package main

import (
	c "../core/crossfire"
	"net"
	"time"
)

func main() {
	service := ":6000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	c.checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	c.checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn) {
	defer conn.Close()
	//业务逻辑
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
