package main

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpctest/services"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	cert,err:=tls.LoadX509KeyPair("server_cert/server.pem","server_cert/server.key")
	if err!=nil{
		log.Fatal("LoadX509KeyPair error:",err)
	}
    certPool:= x509.NewCertPool()
    ca,err:=ioutil.ReadFile("server_cert/ca.pem")
    if err!= nil{
    	log.Fatal("ioutil.ReadFile:",err)
	}
    certPool.AppendCertsFromPEM(ca)

    creds:=credentials.NewTLS(&tls.Config{
    	Certificates: []tls.Certificate{cert},
    	ClientAuth: tls.RequireAndVerifyClientCert,
    	ClientCAs: certPool,
	})
	rpcServer:=grpc.NewServer(grpc.Creds(creds))

	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	listen,_:= net.Listen("tcp",":8080")
	rpcServer.Serve(listen)
}
