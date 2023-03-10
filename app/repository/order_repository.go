package repository

import (
	"context"
	"membership/domain"
	"membership/pb"
)

type OrderRepository struct {
	order pb.OrderServiceClient
}

func NewOrderRepository(order pb.OrderServiceClient) domain.OrderRepository {
	return &OrderRepository{
		order: order,
	}
}

// func (o *OrderRepository) {}

func (o *OrderRepository) SumIncome(ctx context.Context, req *pb.OrderSumIncomeRequest) (res *pb.OrderSumResponse, err error) {
	return o.order.SumIncome(ctx, req)
}

func (o *OrderRepository) Create(ctx context.Context, req *pb.OrderCreateRequest) (*pb.Empty, error) {
	return o.order.Create(ctx, req)
}

func (o *OrderRepository) ChangeStatus(ctx context.Context, req *pb.OrderChangeStatus) (*pb.OperationResponse, error) {
	return o.order.ChangeStatus(ctx, req)
}

func (o *OrderRepository) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (*pb.OrderFindOneResponse, error) {
	return o.order.FindOne(ctx, req)
}

func (o *OrderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	return o.order.FindAll(ctx, req)
}
