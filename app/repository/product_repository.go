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

// withTimeout is a helper function that adds a timeout to a function call.
// It takes in a duration in seconds and a function to be called, and it creates a context
// with a timeout based on the duration. The function is then called with the context.
// If the duration is not specified or is 0, a default timeout of 10 seconds is used.
func (p *ProductRepository) withTimeout(dur int, f func(context.Context)) {
	if dur == 0 {
		dur = 10 // set default timeout duration to 10 seconds
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(dur)*time.Second)
	defer cancel()

	f(ctx)
}

// Post creates a new product in the repository.
func (p *ProductRepository) Post(req *pb.ProductCreateRequest) (err error) {
	p.withTimeout(0, func(ctx context.Context) {
		_, err = p.product.Create(ctx, req)
	})
	return
}

// FindOne retrieves a single product from the repository.
func (p *ProductRepository) FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.FindOne(ctx, req)
	})
	return
}

// FindAll retrieves all products from the repository.
func (p *ProductRepository) FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error) {
	p.withTimeout(15, func(ctx context.Context) {
		res, err = p.product.FindAll(ctx, req)
	})
	return
}

// Update updates a product in the repository.
func (p *ProductRepository) Update(req *pb.ProductUpdateRequest) (res *pb.OperationResponse, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.Update(ctx, req)
	})
	return
}

// Delete removes a product from the repository.
func (p *ProductRepository) Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	p.withTimeout(0, func(ctx context.Context) {
		res, err = p.product.Delete(ctx, req)
	})
	return
}
