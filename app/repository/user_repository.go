package repository

import (
	"context"
	"membership/domain"
	"membership/pb"
)

type userRepository struct {
	user pb.UserServiceClient
}

func NewUserRepository(user pb.UserServiceClient) domain.UserRepository {
	return &userRepository{
		user: user,
	}
}

func (u *userRepository) Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error) {
	res, err = u.user.Find(ctx, &pb.UserFindRequest{UserId: userId})
	return
}
