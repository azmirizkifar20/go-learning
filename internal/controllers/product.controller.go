package controllers

import (
	"strconv"

	"go-learning/internal/dtos"
	"go-learning/internal/helpers"
	"go-learning/internal/models"
	"go-learning/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service  *services.Productservice
	response *helpers.Response
}

func NewProductController(service *services.Productservice, response *helpers.Response) *ProductController {
	return &ProductController{
		service:  service,
		response: response,
	}
}

/*
===========================================================================================
Method: POST
URL Path: /v1/products
===========================================================================================
*/
func (h *ProductController) CreateProduct(c *fiber.Ctx) error {
	// get body
	// var req models.Product
	var req dtos.CreateProductRequest // ambil dari DTO
	if err := c.BodyParser(&req); err != nil {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Invalid body", err.Error())
	}

	// validasi body pakai validator
	if err := helpers.Validate.Struct(req); err != nil {
		formatted := helpers.FormatValidationError(err)
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Validation failed", formatted)
	}

	// mapping DTO ke model
	product := models.Product{
		ProductName: req.ProductName,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
	}

	// create product
	p, err := h.service.Create(c.Context(), &product)
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Internal server error", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, p, "Berhasil membuat product", nil)
}

/*
===========================================================================================
Method: GET
URL Path: /v1/products
===========================================================================================
*/
func (h *ProductController) ListProduct(c *fiber.Ctx) error {
	products, err := h.service.List(c.Context())
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Internal server error", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, products, "Berhasil get data products", nil)
}

/*
===========================================================================================
Method: GET
URL Path: /v1/products/:id
===========================================================================================
*/
func (h *ProductController) GetProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	product, err := h.service.Get(c.Context(), uint(id))
	if err != nil {
		return h.response.Send(c, fiber.StatusNotFound, nil, "product not found", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, product, "Berhasil get data product", nil)
}

/*
===========================================================================================
Method: PUT
URL Path: /v1/products/:id
===========================================================================================
*/
func (h *ProductController) UpdateProduct(c *fiber.Ctx) error {
	// get params id
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	// get body
	var req models.Product
	if err := c.BodyParser(&req); err != nil {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Invalid body!", err.Error())
	}

	_, err := h.service.Get(c.Context(), uint(id))
	if err != nil {
		return h.response.Send(c, fiber.StatusNotFound, nil, "product not found", err.Error())
	}

	// update the product
	product, err := h.service.Update(c.Context(), uint(id), &req)
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Internal server error", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, product, "Success update product!", nil)
}

/*
===========================================================================================
Method: DELETE
URL Path: /v1/products/:id
===========================================================================================
*/
func (h *ProductController) DeleteProduct(c *fiber.Ctx) error {
	// get params id
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.Delete(c.Context(), uint(id))
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Failed delete data", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, nil, "Berhasil hapus product!", nil)
}
