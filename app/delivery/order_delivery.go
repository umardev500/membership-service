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

	router.Post("/orders", handler.Create)
	router.Get("/orders", handler.FindAll)
	router.Get("/orders/:id", handler.FindOne)
	router.Put("/orders/:id", handler.ChangeStatus)
}

func (o *OrderDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

// func (o *OrderDelivery) FindAll(ctx *fiber.Ctx) error {
// reqContext := ctx.Context()
// }

func (o *OrderDelivery) Create(ctx *fiber.Ctx) error {
	var payload = new(pb.OrderCreateRequest)
	if err := ctx.BodyParser(&payload); err != nil {
		return o.handleResponse(ctx, err, 500, "", nil)
	}

	reqContext := ctx.Context()

	products := []*pb.OrderProduct{}
	for _, val := range payload.Product {
		each := pb.OrderProduct{
			ProductId:   val.ProductId,
			Name:        val.Name,
			Price:       val.Price,
			Duration:    val.Duration,
			Description: val.Description,
		}
		products = append(products, &each)
	}

	values := &pb.OrderCreateRequest{
		Buyer: &pb.OrderBuyer{
			CustomerId: payload.Buyer.CustomerId,
			Name:       payload.Buyer.Name,
			User:       payload.Buyer.User,
		},
		Product: products,
	}
	_, err := o.usecase.Create(reqContext, values)
	return o.handleResponse(ctx, err, 200, "Create order", nil)
}

func (o *OrderDelivery) ChangeStatus(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")

	var payload = new(pb.OrderChangeStatus)
	if err := ctx.BodyParser(&payload); err != nil {
		return o.handleResponse(ctx, err, 500, "", nil)
	}

	reqContext := ctx.Context()
	res, err := o.usecase.ChangeStatus(reqContext, &pb.OrderChangeStatus{OrderId: id, Status: payload.Status})

	if !res.IsAffected {
		return o.handleResponse(ctx, err, 304, "Nothing to change", nil)
	}

	return o.handleResponse(ctx, err, 200, "Change status", nil)
}

func (o *OrderDelivery) FindOne(ctx *fiber.Ctx) error {
	var id = ctx.Params("id")
	reqContext := ctx.Context()
	res, err := o.usecase.FindOne(reqContext, &pb.OrderFindOneRequest{OrderId: id})

	return o.handleResponse(ctx, err, 200, "Find order", res)
}

func (o *OrderDelivery) FindAll(ctx *fiber.Ctx) error {
	sort := ctx.Query("sort", "asc")
	page := helper.ToInt(ctx.Query("page", "0"))
	perPage := helper.ToInt(ctx.Query("per_page"))
	search := ctx.Query("search")
	status := ctx.Query("status")

	reqContext := ctx.Context()
	filter := &pb.OrderFindAllRequest{
		Sort:    sort,
		Page:    page,
		PerPage: perPage,
		Search:  search,
		Status:  status,
	}
	res, err := o.usecase.FindAll(reqContext, filter)

	return o.handleResponse(ctx, err, 200, "Find all orders", res)
}
