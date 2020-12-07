package main

import (
	"google.golang.org/grpc"
	"grpctest/services"
	"net"
)

func main() {
	rpcServer:=grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	listen,_:= net.Listen("tcp",":8080")
	rpcServer.Serve(listen)
}
