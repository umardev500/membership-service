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

	router.Get("/products", handler.FindAll)
	router.Post("/products", handler.Create)
	router.Get("/products/:id", handler.FindOne)
	router.Delete("/products/:id", handler.Delete)
}

// func (p *ProductDelivery) Create(ctx *fiber.Ctx) error {}

func (p *ProductDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

func (p *ProductDelivery) FindAll(ctx *fiber.Ctx) error {
	res, err := p.usecase.FindAll(&pb.ProductFindAllRequest{})
	return p.handleResponse(ctx, err, 200, "Find all products", res)
}

func (p *ProductDelivery) FindOne(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	res, err := p.usecase.FindOne(&pb.ProductFindOneRequest{ProductId: id})
	return p.handleResponse(ctx, err, 200, "Find one product", res)
}

func (p *ProductDelivery) Delete(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	res, err := p.usecase.Delete(&pb.ProductDeleteRequest{ProductId: id})
	if err != nil {
		return p.handleResponse(ctx, err, 500, "", nil)
	}

	if !res.IsAffected {
		return p.handleResponse(ctx, err, 304, "Nothing to change", nil)
	}

	return p.handleResponse(ctx, err, 200, "Delete product", nil)
}

func (p *ProductDelivery) Create(ctx *fiber.Ctx) error {
	var payload = new(pb.ProductCreateRequest)

	if err := ctx.BodyParser(&payload); err != nil {
		return p.handleResponse(ctx, err, 500, "", nil)
	}

	product := &pb.ProductCreateRequest{
		Name:        payload.Name,
		Price:       payload.Price,
		Duration:    payload.Duration,
		Description: payload.Description,
	}

	err := p.usecase.Post(product)
	return p.handleResponse(ctx, err, 200, "Create product", nil)
}
