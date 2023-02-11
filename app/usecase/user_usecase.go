package usecase

import (
	"context"
	"membership/domain"
	"membership/pb"
)

type userUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(repository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repository: repository,
	}
}

func (u *userUsecase) Find(ctx context.Context, userId string) (res *pb.User, err error) {
	res, err = u.repository.Find(ctx, userId)
	return
}
