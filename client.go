package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpctest/services"
	"log"
)

func main() {
	conn,err:= grpc.Dial(":8080",grpc.WithInsecure())
	if err!= nil{
		fmt.Println("dial err:",err)
	}

	prodClient:=services.NewProdServiceClient(conn)
	prodRes,err:=prodClient.GetProdStock(context.Background(),&services.ProdRequest{ProdId: 12})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
