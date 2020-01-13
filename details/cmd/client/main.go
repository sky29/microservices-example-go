package main

import (
	"microservices-example-go/details/internal/grpc/domain"
	"microservices-example-go/details/internal/grpc/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7001"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewDetailsServiceClient(conn)

	for i := range [1]int{} {
		detailsModel := domain.Details{
			Id:        int64(i),
			Description:      string("Description: This is a Grpc-Demo"),
		}

		if responseMessage, e := client.Get(context.Background(), &detailsModel); e != nil {
			panic(fmt.Sprintf("Was not able to get product details %v", e))
		} else {
			fmt.Println("Getting Product details:")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
