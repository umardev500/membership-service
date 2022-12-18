package repository

import (
	"context"
	"membership/domain"
	"membership/pb"
	"time"
)

type ProductRepository struct {
	product pb.ProductServiceClient
}

func NewProductRepository(product pb.ProductServiceClient) domain.ProductRepository {
	return &ProductRepository{
		product: product,
	}
}

// func (p *ProductRepository)  {
// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// defer cancel()
// }

func (p *ProductRepository) withTimeout(dur int, f func(context.Context)) {
	if dur == 0 {
		dur = 10 // set default timeout duration to 10 seconds
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dur)*time.Second)
	defer cancel()

	f(ctx)
}

func (p *ProductRepository) Update(req *pb.ProductUpdateRequest) (res *pb.OperationResponse, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.Update(ctx, req)
	})
	return
}

func (p *ProductRepository) FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error) {
	p.withTimeout(15, func(ctx context.Context) {
		res, err = p.product.FindAll(ctx, req)
	})
	return
}

func (p *ProductRepository) FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.FindOne(ctx, req)
	})
	return
}

func (p *ProductRepository) Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.Delete(ctx, req)
	})
	return
}

// Post product data to rpc server
func (p *ProductRepository) Post(req *pb.ProductCreateRequest) (err error) {
	p.withTimeout(0, func(ctx context.Context) {
		_, err = p.product.Create(ctx, req)
	})
	return
}
