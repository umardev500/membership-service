package usecase

import (
	"membership/domain"
	"membership/pb"
)

type ProductUsecase struct {
	repository domain.ProductRepository
}

func NewProductUsecase(repository domain.ProductRepository) domain.ProductUsecase {
	return &ProductUsecase{
		repository: repository,
	}
}

func (p *ProductUsecase) FindOne(req *pb.ProductFindOneRequest) (res *pb.Product, err error) {
	res, err = p.repository.FindOne(req)

	return
}

func (p *ProductUsecase) Delete(req *pb.ProductDeleteRequest) (res *pb.OperationResponse, err error) {
	res, err = p.repository.Delete(req)

	return
}

// Post product data
func (p *ProductUsecase) Post(req *pb.ProductCreateRequest) (err error) {
	err = p.repository.Post(req)

	return
}
