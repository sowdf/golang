package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{A: 2, B: 3}, &result)

	if err != nil {
		log.Printf("error : %s", err)
	} else {
		fmt.Println(result)
	}

}
