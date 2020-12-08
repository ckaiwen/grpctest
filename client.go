package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpctest/services"
	"io/ioutil"
	"log"
)

func main() {
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

	conn,err:= grpc.Dial(":8080",grpc.WithTransportCredentials(creds))
	prodClient:=services.NewProdServiceClient(conn)
	prodRes,err:=prodClient.GetProdStock(context.Background(),&services.ProdRequest{ProdId: 15})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
