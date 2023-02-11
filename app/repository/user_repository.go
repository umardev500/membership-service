package repository

import (
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
