package domain

import (
	"context"
	"membership/pb"
)

type UserLocation struct {
	Address    string `bson:"address" json:"address,omitempty"`
	Village    string `bson:"village" json:"village,omitempty"`
	District   string `bon:"district" json:"district,omitempty"`
	City       string `bson:"city" json:"city,omitempty"`
	Province   string `bson:"province" json:"province,omitempty"`
	PostalCode string `bson:"postal_code" json:"postal_code,omitempty"`
}

type UserDetail struct {
	Name     string        `bson:"name" json:"name,omitempty"`
	Email    string        `bson:"email" json:"email,omitempty"`
	Phone    string        `bson:"phone" json:"phone,omitempty"`
	Avatar   string        `bson:"avatar" json:"avatar,omitempty"`
	Gender   string        `bson:"gender" json:"gender,omitempty"`
	Location *UserLocation `bson:"location" json:"location,omitempty"`
}

type UserUsecase interface {
	Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error)
	UpdateDetail(ctx context.Context, userId string, detail *pb.UserDetail) (res *pb.OperationResponse, err error)
}

type UserRepository interface {
	Find(ctx context.Context, userId string) (res *pb.UserFindResponse, err error)
	UpdateDetail(ctx context.Context, req *pb.UserUpdateDetailRequest) (res *pb.OperationResponse, err error)
}
