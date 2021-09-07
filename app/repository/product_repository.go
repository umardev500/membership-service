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

func (p *ProductRepository) FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = p.product.FindOne(ctx, req)

	return
}

func (p *ProductRepository) Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = p.product.Delete(ctx, req)

	return
}

// Post product data to rpc server
func (p *ProductRepository) Post(req *pb.ProductCreateRequest) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = p.product.Create(ctx, req)

	return
}
