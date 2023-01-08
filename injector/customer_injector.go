package injector

import (
	"membership/app/delivery"
	"membership/app/repository"
	"membership/app/usecase"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

func NewCustomerInjector(customer pb.CustomerServiceClient, router fiber.Router) {
	repository := repository.NewCustomerRepository(customer)
	usecase := usecase.NewCustomerUsecase(repository)
	delivery.NewCustomerDelivery(usecase, router)
}
