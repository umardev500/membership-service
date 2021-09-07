package delivery

import (
	"membership/domain"
	"membership/helper"
	"membership/pb"

	"github.com/gofiber/fiber/v2"
)

type ProductDelivery struct {
	usecase domain.ProductUsecase
}

func NewProductDelivery(usecase domain.ProductUsecase, router fiber.Router) {
	handler := &ProductDelivery{
		usecase: usecase,
	}

	router.Post("/product", handler.Create)
	router.Get("/product/:id", handler.FindOne)
	router.Delete("/product/:id", handler.Delete)
}

// func (p *ProductDelivery) Create(ctx *fiber.Ctx) error {}

func (p *ProductDelivery) FindOne(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	res, err := p.usecase.FindOne(&pb.ProductFindOneRequest{ProductId: id})
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}

	return helper.ApiResponse(ctx, 200, "Delete product", res)
}

func (p *ProductDelivery) Delete(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	res, err := p.usecase.Delete(&pb.ProductDeleteRequest{ProductId: id})
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}

	if !res.IsAffected {
		return helper.ApiResponse(ctx, 304, "Nothing to change", nil)
	}

	return helper.ApiResponse(ctx, 200, "Delete product", nil)
}

func (p *ProductDelivery) Create(ctx *fiber.Ctx) error {
	var payload = new(pb.ProductCreateRequest)

	if err := ctx.BodyParser(&payload); err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}

	product := &pb.ProductCreateRequest{
		Name:        payload.Name,
		Price:       payload.Price,
		Duration:    payload.Duration,
		Description: payload.Description,
	}

	err := p.usecase.Post(product)
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}

	return helper.ApiResponse(ctx, 200, "Create product", nil)
}
