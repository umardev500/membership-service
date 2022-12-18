package domain

import (
	"context"
	"membership/pb"
)

type OrderUsecase interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
}

type OrderRepository interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
}
