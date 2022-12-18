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

func (o *OrderRepository) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	res, err = o.order.FindAll(ctx, req)

	return
}
