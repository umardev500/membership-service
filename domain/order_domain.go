package domain

import (
	"context"
	"membership/pb"
)

type OrderUsecase interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
	Create(context.Context, *pb.OrderCreateRequest) (*pb.Empty, error)
}

type OrderRepository interface {
	Create(context.Context, *pb.OrderCreateRequest) (*pb.Empty, error)
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
}
