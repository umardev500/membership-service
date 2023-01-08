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
}

func (c *CustomerDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
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
