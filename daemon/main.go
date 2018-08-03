package main

import (
	"net"
	"net/rpc"
	"carwer/distribute/rpcdaemon"
	"github.com/emicklei/go-restful/log"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(rpcdaemon.RpcDaemonService{})
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error :%v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
