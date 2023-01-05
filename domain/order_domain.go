package domain

import (
	"context"
	"membership/pb"
)

type OrderRequest struct {
	Buyer struct {
		CustomerId string `json:"customer_id"`
		Name       string `json:"name"`
		User       string `json:"user"`
	} `json:"buyer"`
	Product []struct {
		ProductId   string `json:"product_id"`
		Name        string `json:"name"`
		Price       int64  `json:"price"`
		Duration    int64  `json:"duration"`
		Description string `json:"description"`
	} `json:"product"`
}

type OrderBankPermataRequest struct {
	OrderRequest
	Payment PermataRequest `json:"payment"`
}

type OrderUsecase interface {
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
	Create(ctx context.Context, orderId string, req *pb.OrderCreateRequest) (*pb.Empty, error)
}

type OrderRepository interface {
	Create(ctx context.Context, orderId string, req *pb.OrderCreateRequest) (*pb.Empty, error)
	FindAll(context.Context, *pb.OrderFindAllRequest) (*pb.OrderFindAllResponse, error)
	FindOne(context.Context, *pb.OrderFindOneRequest) (*pb.Order, error)
	ChangeStatus(context.Context, *pb.OrderChangeStatus) (*pb.OperationResponse, error)
}
