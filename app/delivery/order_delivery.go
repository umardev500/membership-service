package delivery

import (
	"membership/domain"
	"membership/helper"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

type OrderDelivery struct {
	usecase domain.OrderUsecase
}

func NewOrderDelivery(usecase domain.OrderUsecase, router fiber.Router) {
	handler := &OrderDelivery{
		usecase: usecase,
	}

	router.Get("/orders", handler.FindAll)
}

func (o *OrderDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

// func (o *OrderDelivery) FindAll(ctx *fiber.Ctx) error {}

func (o *OrderDelivery) FindAll(ctx *fiber.Ctx) error {
	reqContext := ctx.Context()
	res, err := o.usecase.FindAll(reqContext, &pb.OrderFindAllRequest{})

	return o.handleResponse(ctx, err, 200, "Find all orders", res)
}
