package delivery

import (
	"membership/domain"

	"github.com/gofiber/fiber/v2"
)

type userDelivery struct {
	usecase domain.UserUsecase
}

func NewUserDelivery(usecase domain.UserUsecase, router fiber.Router) {
	handler := &userDelivery{
		usecase: usecase,
	}

	router.Get("/user/:id", handler.Find)
}

func (cu *userDelivery) Find(ctx *fiber.Ctx) error {
	return nil
}
