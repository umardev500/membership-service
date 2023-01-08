package domain

import (
	"context"
	"membership/pb"
)

type CustomerUsecase interface {
	FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (*pb.CustomerFindAllResponse, error)
}

type CustomerRepository interface {
	FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (*pb.CustomerFindAllResponse, error)
}
