package usecase

import (
	"context"
	"membership/domain"
	"membership/pb"
	"time"
)

type CustomerUsecase struct {
	repository domain.CustomerRepository
}

func NewCustomerUsecase(repository domain.CustomerRepository) domain.CustomerUsecase {
	return &CustomerUsecase{repository: repository}
}

func (c *CustomerUsecase) withTimeout(dur int64, ctx context.Context, f func(context.Context)) {
	if dur == 0 {
		dur = 10
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(dur)*time.Second)
	defer cancel()

	f(ctx)
}

func (c *CustomerUsecase) FindAll(ctx context.Context, req *pb.CustomerFindAllRequest) (res *pb.CustomerFindAllResponse, err error) {
	c.withTimeout(0, ctx, func(ctx context.Context) {
		res, err = c.repository.FindAll(ctx, req)
	})

	return
}
