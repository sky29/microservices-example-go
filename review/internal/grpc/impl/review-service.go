package impl

import (
	"microservices-example-go/review/internal/grpc/domain"
	"microservices-example-go/review/internal/grpc/service"
	"context"
	"log"
)

type ReviewServiceGrpcImpl struct {
}

func NewReviewServiceGrpcImpl() *ReviewServiceGrpcImpl {
	return &ReviewServiceGrpcImpl{}
}

func (serviceImpl *ReviewServiceGrpcImpl) Get(ctx context.Context, in *domain.Review) (*service.GetReviewResponse, error) {
	log.Println("Received request for getting reviews ")

	return &service.GetReviewResponse{
		Review: in,
	}, nil
}
