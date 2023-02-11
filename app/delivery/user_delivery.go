package delivery

import (
	"membership/domain"
	"membership/helper"

	"github.com/gofiber/fiber/v2"
)

type userDelivery struct {
	usecase domain.UserUsecase
}

func NewUserDelivery(usecase domain.UserUsecase, router fiber.Router) {
	handler := &userDelivery{
		usecase: usecase,
	}

	router.Get("/users/:id", handler.Find)
}

func (u *userDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

func (u *userDelivery) Find(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	reqCtx := ctx.Context()
	res, err := u.usecase.Find(reqCtx, userId)

	return u.handleResponse(ctx, err, 200, "Find user", res)
}
