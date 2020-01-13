package main

import (
	"microservices-example-go/details/internal/grpc/impl"
	"microservices-example-go/details/internal/grpc/service"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7001)
	grpcServer := grpc.NewServer()

	detailsServiceImpl := impl.NewDetailsServiceGrpcImpl()
	service.RegisterDetailsServiceServer(grpcServer, detailsServiceImpl)

	// start the server
	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}
	return lis
}
