package domain

import (
	"context"
	"membership/pb"
)

type OrderUsecase interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
}

type OrderRepository interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
}
