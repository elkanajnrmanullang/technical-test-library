package buku

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/buku/buku_response"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ports.IBukuService
}

func NewBukuController(service ports.IBukuService) *Controller {
	return &Controller{
		service,
	}
}

func (c *Controller) GetAllBuku(f *fiber.Ctx) error {
	bukus, err := c.service.GetAllBuku()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get all buku",
		"data":  buku_response.FromDomainListBuku(bukus),
	})
}

func (c *Controller) GetBukuByID(f *fiber.Ctx) error {
	id := html.EscapeString(strings.TrimSpace(f.Params("id")))
	if id == "" {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "ID is required",
		})
	}

	buku, err := c.service.GetBukuByID(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get buku by ID",
		"data":  buku_response.FromDomainBuku(buku),
	})
}
