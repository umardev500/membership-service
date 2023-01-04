package injector

import (
	"fmt"
	"membership/app/delivery"
	"membership/app/repository"
	"membership/app/usecase"
	"membership/pb"
	"os"

	"github.com/gofiber/fiber/v2"
)

func NewOrderInjector(order pb.OrderServiceClient, router fiber.Router) {
	authStr := "Basic U0ItTWlkLXNlcnZlci1Vbmo2LU5yMVlyZlhBWENYMllSWUxmZXk6"
	chargeURL := fmt.Sprintf("%s/%s", os.Getenv("PG_URL"), "charge")
	pgRepo := repository.NewPGRepository(chargeURL, authStr)
	pgUsecase := usecase.NewPGUsecase(pgRepo)

	repository := repository.NewOrderRepository(order)
	usecase := usecase.NewOrderUsecase(repository)

	delivery.NewOrderDelivery(usecase, pgUsecase, router)
}
