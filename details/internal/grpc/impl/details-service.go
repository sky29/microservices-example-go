package impl

import (
	"microservices-example-go/details/internal/grpc/domain"
	"microservices-example-go/details/internal/grpc/service"
	"context"
	"log"
)

type DetailsServiceGrpcImpl struct {
}

func NewDetailsServiceGrpcImpl() *DetailsServiceGrpcImpl {
	return &DetailsServiceGrpcImpl{}
}

func (serviceImpl *DetailsServiceGrpcImpl) Get(ctx context.Context, in *domain.Details) (*service.GetDetailsResponse, error) {
	log.Println("Received request for getting product details ")

	return &service.GetDetailsResponse{
		Details: in,
	}, nil
}
