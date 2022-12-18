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

// func (p *ProductUsecase) {}
func (p *ProductUsecase) FindAll(req *pb.ProductFindAllRequest) (*pb.ProductFindAllResponse, error) {
	return p.repository.FindAll(req)
}

func (p *ProductUsecase) FindOne(req *pb.ProductFindOneRequest) (*pb.Product, error) {
	return p.repository.FindOne(req)
}

func (p *ProductUsecase) Delete(req *pb.ProductDeleteRequest) (*pb.OperationResponse, error) {
	return p.repository.Delete(req)
}

func (p *ProductUsecase) Post(req *pb.ProductCreateRequest) error {
	return p.repository.Post(req)
}
