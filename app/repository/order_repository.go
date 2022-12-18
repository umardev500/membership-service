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

func (o *OrderRepository) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (*pb.Order, error) {
	return o.order.FindOne(ctx, req)
}

func (o *OrderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	return o.order.FindAll(ctx, req)
}
