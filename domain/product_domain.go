package domain

import "membership/pb"

type ProductUsecase interface {
	Post(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
}

type ProductRepository interface {
	Post(req *pb.ProductCreateRequest) error
	Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error)
}
