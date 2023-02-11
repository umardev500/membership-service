package usecase

import (
	"membership/domain"
)

type userUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repository: repository,
	}
}
