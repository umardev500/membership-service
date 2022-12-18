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

func (p *ProductRepository) withTimeout(f func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return f(ctx)
}

func (p *ProductRepository) Update(req *pb.ProductUpdateRequest) (res *pb.OperationResponse, err error) {
	err = p.withTimeout(func(ctx context.Context) error {
		res, err = p.product.Update(ctx, req)
		return err
	})
	return
}

func (p *ProductRepository) FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error) {
	err = p.withTimeout(func(ctx context.Context) error {
		res, err = p.product.FindAll(ctx, req)
		return err
	})
	return
}

func (p *ProductRepository) FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error) {
	err = p.withTimeout(func(ctx context.Context) error {
		res, err = p.product.FindOne(ctx, req)
		return err
	})
	return
}

func (p *ProductRepository) Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	err = p.withTimeout(func(ctx context.Context) error {
		res, err = p.product.Delete(ctx, req)
		return err
	})
	return
}

// Post product data to rpc server
func (p *ProductRepository) Post(req *pb.ProductCreateRequest) (err error) {
	err = p.withTimeout(func(ctx context.Context) error {
		_, err = p.product.Create(ctx, req)
		return err
	})
	return
}
