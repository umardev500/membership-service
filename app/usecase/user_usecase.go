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

func (u *userUsecase) UpdateCreds(ctx context.Context, req *pb.UserUpdateCredsRequest) (res *pb.OperationResponse, err error) {
	return u.repository.UpdateCreds(ctx, req)
}

func (u *userUsecase) UpdateDetail(ctx context.Context, userId string, detail *pb.UserDetail) (res *pb.OperationResponse, err error) {
	payload := &pb.UserUpdateDetailRequest{UserId: userId, Detail: detail}
	return u.repository.UpdateDetail(ctx, payload)
}

func (u *userUsecase) Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error) {
	res, err = u.repository.Find(ctx, userId)
	return
}
