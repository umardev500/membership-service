package delivery

import (
	"membership/domain"
	"membership/helper"
	"membership/pb"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerDelivery struct {
	usecase domain.CustomerUsecase
}

func NewCustomerDelivery(usecase domain.CustomerUsecase, router fiber.Router) {
	handler := &CustomerDelivery{usecase: usecase}
	router.Get("/customers", handler.FindAll)
	router.Delete("/customers/:id", handler.Delete)
}

func (c *CustomerDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

func (c *CustomerDelivery) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	hard, _ := strconv.ParseBool(ctx.Query("hard", "false"))

	reqCtx := ctx.Context()
	res, err := c.usecase.Delete(reqCtx, &pb.CustomerDeleteRequest{CustomerId: id, Hard: hard})

	return c.handleResponse(ctx, err, 200, "Delete customer", res)
}

func (c *CustomerDelivery) FindAll(ctx *fiber.Ctx) error {
	sort := ctx.Query("sort", "asc")
	page := helper.ToInt(ctx.Query("page", "0"))
	perPage := helper.ToInt(ctx.Query("per_page"))
	search := ctx.Query("search")
	status := ctx.Query("status")
	countOnly, _ := strconv.ParseBool(ctx.Query("count_only"))

	reqCtx := ctx.Context()
	filter := &pb.CustomerFindAllRequest{
		Sort:      sort,
		Page:      page,
		PerPage:   perPage,
		Search:    search,
		Status:    status,
		CountOnly: countOnly,
	}
	res, err := c.usecase.FindAll(reqCtx, filter)

	return c.handleResponse(ctx, err, 200, "Find all customers", res)
}
