package repository

import (
	"context"
	"membership/domain"
	"membership/pb"
)

type CustomerRepository struct {
	customer pb.CustomerServiceClient
}

func NewCustomerRepository(customer pb.CustomerServiceClient) domain.CustomerRepository {
	return &CustomerRepository{customer: customer}
}

func (c *CustomerRepository) FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (*pb.CustomerFindAllResponse, error) {
	return c.customer.FindAll(ctx, req)
}
