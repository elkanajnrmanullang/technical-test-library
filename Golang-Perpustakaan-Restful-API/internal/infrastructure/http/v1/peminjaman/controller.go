package peminjaman

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/peminjaman/peminjaman_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/peminjaman/peminjaman_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ports.IPeminjamanService
}

func NewPeminjamanController(service ports.IPeminjamanService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetAllPeminjaman(f *fiber.Ctx) error {
	peminjamen, err := c.service.GetAllPeminjaman()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get all peminjaman",
		"data":  peminjaman_response.FromDomainList(peminjamen),
	})
}

func (c *Controller) GetPeminjamanByID(f *fiber.Ctx) error {
	id := html.EscapeString(strings.TrimSpace(f.Params("id")))
	if id == "" {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "ID is required",
		})
	}

	peminjaman, err := c.service.GetPeminjamanByID(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get peminjaman by ID",
		"data":  peminjaman_response.FromDomain(peminjaman),
	})
}

func (c *Controller) GetDetailPeminjaman(f *fiber.Ctx) error {
	id := html.EscapeString(strings.TrimSpace(f.Params("id")))
	if id == "" {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "ID is required",
		})
	}

	peminjaman, err := c.service.GetDetailPeminjaman(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get detail peminjaman",
		"data":  peminjaman_response.FromDomainDetail(peminjaman),
	})
}

func (c *Controller) CreatePeminjaman(f *fiber.Ctx) error {
	req := new(peminjaman_request.CreatePeminjamanRequest)
	if err := f.BodyParser(&req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := utility.GetValidator().Struct(req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.service.CreatePeminjaman(req.ToDomain()); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Peminjaman created successfully",
	})
}

func (c *Controller) UpdatePeminjaman(f *fiber.Ctx) error {
	req := new(peminjaman_request.UpdatePeminjamanRequest)
	if err := f.BodyParser(&req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := utility.GetValidator().Struct(req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.service.UpdatePeminjaman(req.ToDomain()); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Peminjaman updated successfully",
	})
}

func (c *Controller) DeletePeminjaman(f *fiber.Ctx) error {
	req := new(peminjaman_request.DeletePeminjamanRequest)
	if err := f.BodyParser(&req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := utility.GetValidator().Struct(req); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := c.service.DeletePeminjaman(req.IDPeminjaman); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Peminjaman deleted successfully",
	})
}
