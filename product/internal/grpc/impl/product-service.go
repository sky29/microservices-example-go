package impl

import (
	"microservices-example-go/product/internal/grpc/domain"
	"microservices-example-go/product/internal/grpc/service"
	"context"
	"log"
)

type ProductServiceGrpcImpl struct {
}

func NewProductServiceGrpcImpl() *ProductServiceGrpcImpl {
	return &ProductServiceGrpcImpl{}
}

func (serviceImpl *ProductServiceGrpcImpl) Get(ctx context.Context, in *domain.Product) (*service.GetProductResponse, error) {
	log.Println("Received request for getting product page ")

	return &service.GetProductResponse{
		Product: in,
	}, nil
}
