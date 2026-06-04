package user

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
)

type Controller struct {
	service ports.IUserService
}

func NewUserController(service ports.IUserService) *Controller {
	return &Controller{
		service,
	}
}
