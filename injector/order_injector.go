package injector

import (
	"membership/app/delivery"
	"membership/app/repository"
	"membership/app/usecase"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

func NewOrderInjector(order pb.OrderServiceClient, router fiber.Router) {
	repository := repository.NewOrderRepository(order)
	usecase := usecase.NewOrderUsecase(repository)

	delivery.NewOrderDelivery(usecase, router)
}
