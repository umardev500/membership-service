package delivery

import (
	"membership/domain"

	"github.com/gofiber/fiber/v2"
)

type CustomerDelivery struct {
	usecase domain.CustomerUsecase
}

func NewCustomerDelivery(usecase domain.CustomerUsecase, router fiber.Router) {
	handler := &CustomerDelivery{usecase: usecase}
	router.Get("/customers", handler.FindAll)
}

func (c *CustomerDelivery) FindAll(ctx *fiber.Ctx) error {
	return ctx.JSON("Find all")
}
