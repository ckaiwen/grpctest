package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpctest/services"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	cert,err:=tls.LoadX509KeyPair("client_cert/client.pem","client_cert/client.key")
	if err!=nil{
		log.Fatal("LoadX509KeyPair error:",err)
	}
	certPool:= x509.NewCertPool()
	ca,err:=ioutil.ReadFile("client_cert/ca.pem")
	if err!= nil{
		log.Fatal("ioutil.ReadFile:",err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds:=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName: "localhost",
		RootCAs: certPool,
	})

	gwmux:=runtime.NewServeMux()
	opt:=[]grpc.DialOption{grpc.WithTransportCredentials(creds)}
	err=services.RegisterProdServiceHandlerFromEndpoint(context.Background(),gwmux,"localhost:8080",opt)
	if err!= nil{
		log.Fatal()
	}

	httpServer:=&http.Server{
		Addr: ":8081",
		Handler: gwmux,
	}
    httpServer.ListenAndServe()
}
