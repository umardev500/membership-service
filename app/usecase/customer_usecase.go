package usecase

import "membership/domain"

type CustomerUsecase struct {
	repository domain.CustomerRepository
}

func NewCustomerUsecase(repository domain.CustomerRepository) domain.CustomerUsecase {
	return &CustomerUsecase{repository: repository}
}
