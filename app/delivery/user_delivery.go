package delivery

import (
	"fmt"
	"io"
	"membership/domain"
	"membership/helper"
	"membership/pb"
	"membership/variable"
	"net/http"
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
	router.Put("/users/:id/avatar", handler.UpdateAvatar)
	router.Put("/users/:id/detail", handler.UpdateDetail)
	router.Put("/users/:id/location", handler.UpdateLocation)
	router.Put("/users/:id/creds", handler.UpdateCreds)
}

func (u *userDelivery) handleResponse(ctx *fiber.Ctx, err error, status int, message string, data interface{}) error {
	if err != nil {
		return helper.ApiResponse(ctx, 500, err.Error(), nil)
	}
	return helper.ApiResponse(ctx, status, message, data)
}

func (u *userDelivery) UpdateCreds(ctx *fiber.Ctx) error {
	var userData domain.UserCreds
	if err := ctx.BodyParser(&userData); err != nil {
		return u.handleResponse(ctx, err, 200, "Update detail", nil)
	}

	reqCtx := ctx.Context()
	userId := ctx.Params("id")
	payload := &pb.UserUpdateCredsRequest{
		UserId:  userId,
		User:    userData.User,
		Pass:    userData.Pass,
		NewPass: userData.NewPass,
	}

	resp, err := u.usecase.UpdateCreds(reqCtx, payload)
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Update creds", nil)
	}

	return u.handleResponse(ctx, err, 200, "Update creds", resp)
}

func (u *userDelivery) UpdateLocation(ctx *fiber.Ctx) error {
	var userDetail domain.UserDetail
	if err := ctx.BodyParser(&userDetail); err != nil {
		return u.handleResponse(ctx, err, 200, "Update detail", nil)
	}
	detail := &pb.UserDetail{
		Location: &pb.UserLocation{
			Address:    userDetail.Location.Address,
			Village:    userDetail.Location.Village,
			District:   userDetail.Location.District,
			City:       userDetail.Location.City,
			Province:   userDetail.Location.Province,
			PostalCode: userDetail.Location.PostalCode,
		},
	}

	reqCtx := ctx.Context()
	userId := ctx.Params("id")

	resp, err := u.usecase.UpdateDetail(reqCtx, userId, detail)
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Update detail", nil)
	}

	return u.handleResponse(ctx, err, 200, "Update detail", resp)
}

func (u *userDelivery) UpdateDetail(ctx *fiber.Ctx) error {
	var userDetail domain.UserDetail
	if err := ctx.BodyParser(&userDetail); err != nil {
		return u.handleResponse(ctx, err, 200, "Update detail", nil)
	}
	detail := &pb.UserDetail{
		Name:   userDetail.Name,
		Email:  userDetail.Email,
		Phone:  userDetail.Phone,
		Avatar: userDetail.Avatar,
		Gender: userDetail.Gender,
	}

	reqCtx := ctx.Context()
	userId := ctx.Params("id")

	resp, err := u.usecase.UpdateDetail(reqCtx, userId, detail)
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Update detail", nil)
	}

	return u.handleResponse(ctx, err, 200, "Update detail", resp)
}

func (u *userDelivery) UpdateAvatar(ctx *fiber.Ctx) error {
	var responses variable.OpertionResponse = variable.OpertionResponse{IsAffected: true}
	userId := ctx.Params("id")
	reqCtx := ctx.Context()

	// Upload file
	file, err := ctx.FormFile("file")
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	rootPath := "./public"
	fileLocation := fmt.Sprintf("/uploads/images/avatars/%s", file.Filename)
	filePath := fmt.Sprintf("%s%s", rootPath, fileLocation)
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

	detail := &pb.UserDetail{Avatar: fileLocation}
	resp, err := u.usecase.UpdateDetail(reqCtx, userId, detail)
	if err != nil {
		return u.handleResponse(ctx, err, 200, "Upload file", nil)
	}

	if !resp.IsAffected {
		responses = variable.OpertionResponse{IsAffected: false}
		return u.handleResponse(ctx, err, http.StatusNotModified, "Update avatar", responses)
	}

	return u.handleResponse(ctx, err, 200, "Update avatar", responses)
}

func (u *userDelivery) Find(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	reqCtx := ctx.Context()
	res, err := u.usecase.Find(reqCtx, userId)

	return u.handleResponse(ctx, err, 200, "Find user", res)
}
