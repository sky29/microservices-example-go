package main

import (
	"microservices-example-go/product/internal/grpc/domain"
	"microservices-example-go/product/internal/grpc/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7000"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewProductServiceClient(conn)

	for i := range [1]int{} {
		productModel := domain.Product{
			Id:        int64(i),
			Name:      string("Grpc-Demo"),
		}

		if responseMessage, e := client.Get(context.Background(), &productModel); e != nil {
			panic(fmt.Sprintf("Was not able to get product %v", e))
		} else {
			fmt.Println("Getting Product:")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
