package rpcsupport

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func CreateRpcClient(address string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	client := jsonrpc.NewClient(conn)
	return client, nil
}
