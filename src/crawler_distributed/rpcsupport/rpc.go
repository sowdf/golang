package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServiceRpc(host string, server interface{}) error {
	//注册、
	rpc.Register(server)

	//打开tcp

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != err {
			log.Printf("conn fail : %s", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
	return err
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
