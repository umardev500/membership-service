package domain

import (
	"context"
	"membership/pb"
)

type UserUsecase interface {
	Find(ctx context.Context, userId string) (res *pb.User, err error)
}

type UserRepository interface {
	Find(ctx context.Context, userId string) (res *pb.User, err error)
}
