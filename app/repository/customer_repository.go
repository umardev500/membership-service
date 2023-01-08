package repository

import (
	"membership/domain"
	"membership/pb"
)

type CustomerRepository struct {
	customer pb.CustomerServiceClient
}

func NewCustomerRepository(customer pb.CustomerServiceClient) domain.CustomerRepository {
	return &CustomerRepository{customer: customer}
}
