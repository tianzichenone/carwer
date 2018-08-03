package rpcsupport

import (
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
	"log"
)

func CreateRpcServer(service interface{}, address string) error{
	err := rpc.Register(service)
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Printf("running rpc server in address:%s", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error :%v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}