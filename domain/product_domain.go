package domain

import "membership/pb"

type ProductUsecase interface {
	Post(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
	FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error)
	FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error)
}

type ProductRepository interface {
	Post(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
	FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error)
	FindAll(req *pb.ProductFindAllRequest) (res *pb.ProductFindAllResponse, err error)
}
