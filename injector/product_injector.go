package injector

import (
	"membership/app/delivery"
	"membership/app/repository"
	"membership/app/usecase"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

func NewProductInjector(client pb.ProductServiceClient, router fiber.Router) {
	repository := repository.NewProductRepository(client)
	usecase := usecase.NewProductUsecase(repository)

	delivery.NewProductDelivery(usecase, router)
}
