package main

import (
	"net"
	"net/rpc/jsonrpc"
	"carwer/distribute/rpcdaemon"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("RpcDaemonService.Divide", rpcdaemon.Args{A :3, B: 0}, &result)
	fmt.Println(result, err)
}
