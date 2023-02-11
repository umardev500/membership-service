package injector

import (
	"membership/app/delivery"
	"membership/app/repository"
	"membership/app/usecase"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

func NewUserInjector(user pb.UserServiceClient, router fiber.Router) {
	rep := repository.NewUserRepository(user)
	uc := usecase.NewUserUsecase(rep)
	delivery.NewUserDelivery(uc, router)
}
