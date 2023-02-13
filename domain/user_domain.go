package domain

import (
	"context"
	"membership/pb"
)

type UserLocation struct {
	Address    string `bson:"address,omitempty"`
	Village    string `bson:"village,omitempty"`
	District   string `bon:"district,omitempty"`
	City       string `bson:"city,omitempty"`
	Province   string `bson:"province,omitempty"`
	PostalCode string `bson:"postal_code,omitempty"`
}

type UserDetail struct {
	Name     string        `bson:"name,omitempty"`
	Email    string        `bson:"email,omitempty"`
	Phone    string        `bson:"phone,omitempty"`
	Location *UserLocation `bson:"location,omitempty"`
}

type UserUsecase interface {
	Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error)
	UpdateDetail(ctx context.Context, userId string, detail *pb.UserDetail) (res *pb.OperationResponse, err error)
}

type UserRepository interface {
	Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error)
	UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error)
}
