package delivery

import (
	"errors"
	"membership/domain"
	"membership/helper"
	"membership/pb"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderDelivery struct {
	usecase   domain.OrderUsecase
	pgUsecase domain.PGUsecase
}

func NewOrderDelivery(usecase domain.OrderUsecase, pgUsecase domain.PGUsecase, router fiber.Router) {
	handler := &OrderDelivery{
		usecase:   usecase,
		pgUsecase: pgUsecase,
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

func (o *OrderDelivery) createBankPayment(ctx *fiber.Ctx, orderId string) (response *domain.BankPaymentResponse, err error) {
	var payload interface{}
	if err = ctx.BodyParser(&payload); err != nil {
		return
	}

	payloadMap := payload.(map[string]interface{})
	var payment map[string]interface{}
	if payloadMap["payment"] != nil {
		payment = payloadMap["payment"].(map[string]interface{})
	}

	if payment != nil {
		bankTransfer := payment["bank_transfer"].(map[string]interface{})
		bank := bankTransfer["bank"]

		if bank == "permata" {

			resp, err := o.pgUsecase.BankCharge(bank.(string), orderId, payment)
			if err != nil {
				return nil, err
			}
			respCast := resp.(domain.PermataResponse)

			response = &domain.BankPaymentResponse{
				OrderId:     orderId,
				PaymentType: respCast.PaymentType,
				Bank:        "permata",
				VaNumber:    respCast.PermataVaNumber,
				GrossAmount: helper.RemovePriceDot(respCast.GrossAmount),
			}

			statusCode, err := strconv.Atoi(respCast.StatusCode)
			if !(statusCode < 200 || statusCode > 300 || statusCode == 300) {
				return nil, err
			} else {
				err = errors.New(respCast.StatusMessage)
				return nil, err
			}
		}
	}

	return response, nil
}

func (o *OrderDelivery) Create(ctx *fiber.Ctx) error {
	paymentType := ctx.Query("payment_type")
	var payment *pb.OrderPayment
	var paymentResponse *domain.BankPaymentResponse

	var payload = new(domain.OrderBankPermataRequest)
	if err := ctx.BodyParser(&payload); err != nil {
		return o.handleResponse(ctx, err, 500, "", nil)
	}
	orderId := strconv.Itoa(int(time.Now().UTC().UnixNano()))

	if paymentType == "bank" {
		pgResp, err := o.createBankPayment(ctx, orderId)
		paymentResponse = pgResp
		if err != nil {
			return o.handleResponse(ctx, err, 500, "", nil)
		}
	}

	if paymentResponse != nil {
		payment = &pb.OrderPayment{
			PaymentType: paymentResponse.PaymentType,
			OrderId:     paymentResponse.OrderId,
			Bank:        paymentResponse.Bank,
			VaNumber:    paymentResponse.VaNumber,
			GrossAmount: paymentResponse.GrossAmount,
		}
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
		Payment: payment,
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
	countOnly, _ := strconv.ParseBool(ctx.Query("count_only"))

	reqContext := ctx.Context()
	filter := &pb.OrderFindAllRequest{
		Sort:      sort,
		Page:      page,
		PerPage:   perPage,
		Search:    search,
		Status:    status,
		CountOnly: countOnly,
	}
	res, err := o.usecase.FindAll(reqContext, filter)

	return o.handleResponse(ctx, err, 200, "Find all orders", res)
}
