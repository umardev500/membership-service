package domain

import (
	"context"
	"membership/pb"
)

type CustomerUsecase interface {
	Delete(ctx context.Context, req *pb.CustomerDeleteRequest) (*pb.OperationResponse, error)
	FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (*pb.CustomerFindAllResponse, error)
}

type CustomerRepository interface {
	Delete(ctx context.Context, req *pb.CustomerDeleteRequest) (*pb.OperationResponse, error)
	FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (*pb.CustomerFindAllResponse, error)
}
