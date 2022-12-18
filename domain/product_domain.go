package domain

import "membership/pb"

// ProductUsecase is an interface that defines the methods for interacting with products.
type ProductUsecase interface {
	// Post creates a new product.
	Post(req *pb.ProductCreateRequest) error
	// Delete removes a product from the repository.
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
	// FindOne retrieves a single product from the repository.
	FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error)
	// FindAll retrieves all products from the repository.
	FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error)
}

// ProductRepository is an interface that defines the methods for interacting with a repository of products.
type ProductRepository interface {
	// Post creates a new product in the repository.
	Post(req *pb.ProductCreateRequest) error
	// FindOne retrieves a single product from the repository.
	FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error)
	// FindAll retrieves all products from the repository.
	FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error)
	// Update updates a product in the repository.
	Update(req *pb.ProductUpdateRequest) (res *pb.OperationResponse, err error)
	// Delete removes a product from the repository.
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
}
