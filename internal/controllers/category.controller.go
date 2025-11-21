package controllers

import (
	"go-learning/internal/helpers"
	"go-learning/internal/services"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	service  *services.CategoryService
	response *helpers.Response
}

func NewCategoryController(service *services.CategoryService, response *helpers.Response) *CategoryController {
	return &CategoryController{
		service:  service,
		response: response,
	}
}

/*
===========================================================================================
Method: POST
URL Path: /v1/categories
===========================================================================================
*/
func (h *CategoryController) CreateCategory(c *fiber.Ctx) error {

	// ambil body
	name := c.FormValue("category_name")
	if name == "" {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "category_name is required", nil)
	}

	// ambil file image
	file, err := c.FormFile("image")
	if err != nil && err != fiber.ErrUnprocessableEntity {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Invalid image file", nil)
	}

	if err == fiber.ErrUnprocessableEntity {
		file = nil
	}

	category, err := h.service.CreateCategory(c.Context(), name, file)
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Interval server error", err.Error())
	}

	return h.response.Send(c, fiber.StatusCreated, category, "Success create product category", nil)
}
