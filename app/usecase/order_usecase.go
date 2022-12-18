package usecase

import (
	"context"
	"membership/domain"
	"membership/pb"
	"time"
)

type OrderUsecase struct {
	repository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &OrderUsecase{
		repository: repository,
	}
}

func (o *OrderUsecase) withTimeout(dur int, ctx context.Context, f func(context.Context)) {
	if dur == 0 {
		dur = 10
	}

	ctx, cancel := context.WithTimeout(ctx, time.Duration(dur)*time.Second)
	defer cancel()

	f(ctx)
}

// func (o *OrderUsecase) {}
func (o *OrderUsecase) FindOne(ctx context.Context, req *pb.OrderFindOneRequest) (res *pb.Order, err error) {
	o.withTimeout(0, ctx, func(ctx context.Context) {
		res, err = o.repository.FindOne(ctx, req)
	})

	return
}

func (o *OrderUsecase) FindAll(ctx context.Context, req *pb.OrderFindAllRequest) (res *pb.OrderFindAllResponse, err error) {
	o.withTimeout(0, ctx, func(ctx context.Context) {
		res, err = o.repository.FindAll(ctx, req)
	})

	return
}
