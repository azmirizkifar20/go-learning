package controllers

import (
	"strconv"

	"go-learning/internal/database"
	"go-learning/internal/helpers"
	"go-learning/internal/models"
	"go-learning/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service  services.ProductService
	response *helpers.Response
}

func NewProductController(service services.ProductService, response *helpers.Response) *ProductController {
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
	db := database.GetDB()

	// get body
	var req models.Product
	if err := c.BodyParser(&req); err != nil {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Invalid body", err.Error())
	}

	// create product
	p, err := h.service.Create(db, &req)
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Internal server error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

/*
===========================================================================================
Method: GET
URL Path: /v1/products
===========================================================================================
*/
func (h *ProductController) ListProduct(c *fiber.Ctx) error {
	db := database.GetDB()
	products, err := h.service.List(db)
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
	db := database.GetDB()
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	product, err := h.service.Get(db, uint(id))
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
	db := database.GetDB()

	// get params id
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	// get body
	var req models.Product
	if err := c.BodyParser(&req); err != nil {
		return h.response.Send(c, fiber.StatusBadRequest, nil, "Invalid body!", err.Error())
	}

	_, err := h.service.Get(db, uint(id))
	if err != nil {
		return h.response.Send(c, fiber.StatusNotFound, nil, "product not found", err.Error())
	}

	// update the product
	product, err := h.service.Update(db, uint(id), &req)
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
	db := database.GetDB()

	// get params id
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.Delete(db, uint(id))
	if err != nil {
		return h.response.Send(c, fiber.StatusInternalServerError, nil, "Failed delete data", err.Error())
	}

	return h.response.Send(c, fiber.StatusOK, nil, "Berhasil hapus product!", nil)
}
