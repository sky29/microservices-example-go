package main

import (
	"microservices-example-go/review/internal/grpc/domain"
	"microservices-example-go/review/internal/grpc/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "localhost:7002"

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := service.NewReviewServiceClient(conn)

	for i := range [1]int{} {
		reviewModel := domain.Review{
			Id:        int64(i),
			Review:      string("Review: This is awesome demo of Grpc"),
			Rating: 	 string ("Rating: *****"),
		}

		if responseMessage, e := client.Get(context.Background(), &reviewModel); e != nil {
			panic(fmt.Sprintf("Was not able to get review %v", e))
		} else {
			fmt.Println("Getting Review:")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
