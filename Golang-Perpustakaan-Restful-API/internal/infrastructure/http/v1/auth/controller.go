package auth

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user/user_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ports.IAuthService
}

func NewAuthController(service ports.IAuthService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) Login(f *fiber.Ctx) error {
	loginReq := new(user_request.AnggotaLoginRequest)
	if err := f.BodyParser(&loginReq); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	cekValid := utility.GetValidator().Struct(loginReq)
	if cekValid != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   user_response.ErrInvalidFormatJson,
		})
	}

	res, err := c.service.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		return f.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Login successful",
		"data":  res,
	})
}
