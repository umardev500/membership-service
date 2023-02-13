package delivery

import (
	"encoding/json"
	"fmt"
	"io"
	"membership/domain"
	"membership/helper"
	"os"

	"github.com/gofiber/fiber/v2"
)

type userDelivery struct {
	usecase domain.UserUsecase
}

func NewUserDelivery(usecase domain.UserUsecase, router fiber.Router) {
	handler := &userDelivery{
		usecase: usecase,
	}

	router.Get("/users/:id", handler.Find)
	router.Put("/users/:id", handler.UpdateDetail)
}

func (u *userDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

func (u *userDelivery) UpdateDetail(ctx *fiber.Ctx) error {
	// Read payload
	var userDetail domain.UserDetail

	req, _ := ctx.MultipartForm()
	payload := req.Value["payload"][0]
	json.Unmarshal([]byte(payload), &userDetail)

	// Upload file
	file, err := ctx.FormFile("file")
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	filePath := fmt.Sprintf("./public/uploads/images/avatars/%s", file.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	defer out.Close()

	fileReader, err := file.Open()
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	defer fileReader.Close()

	var contents []byte
	var buf = make([]byte, 1024)
	for {
		n, err := fileReader.Read(buf)
		if err != nil && err != io.EOF {
			return u.handleResponse(ctx, err, 200, "Upload file", nil)
		}

		if n == 0 {
			break
		}

		contents = append(contents, buf[:n]...)

	}

	if _, err := out.Write(contents); err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	return ctx.JSON("OK")
}

func (u *userDelivery) Find(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	reqCtx := ctx.Context()
	res, err := u.usecase.Find(reqCtx, userId)

	return u.handleResponse(ctx, err, 200, "Find user", res)
}
