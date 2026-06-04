package denda

import (
	"html"
	"strings"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/denda/denda_request"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/denda/denda_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ports.IDendaService
}

func NewDendaController(service ports.IDendaService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) GetAllDenda(f *fiber.Ctx) error {
	dendas, err := c.service.GetAllDenda()
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get all denda",
		"data":  denda_response.FromDomainList(dendas),
	})
}

func (c *Controller) GetDendaByID(f *fiber.Ctx) error {
	id := html.EscapeString(strings.TrimSpace(f.Params("id")))
	if id == "" {
		return f.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "ID is required",
		})
	}

	denda, err := c.service.GetDendaByID(id)
	if err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Success get denda by ID",
		"data":  denda_response.FromDomain(denda),
	})
}

func (c *Controller) CreateDenda(f *fiber.Ctx) error {
	req := new(denda_request.CreateDendaRequest)
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

	if err := c.service.CreateDenda(req.ToDomain()); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   "Denda created successfully",
	})
}

func (c *Controller) UpdateDenda(f *fiber.Ctx) error {
	req := new(denda_request.UpdateDendaRequest)
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

	if err := c.service.UpdateDenda(req.ToDomain()); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Denda updated successfully",
	})
}

func (c *Controller) DeleteDenda(f *fiber.Ctx) error {
	req := new(denda_request.DeleteDendaRequest)
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

	if err := c.service.DeleteDenda(req.IDDenda); err != nil {
		return f.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return f.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Denda deleted successfully",
	})
}
