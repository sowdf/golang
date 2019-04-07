package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc"
)

func main() {
	//注册、
	rpc.Register(rpcdemo.DemoService{})

	//打开tcp

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != err {
			log.Printf("conn fail : %s", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
