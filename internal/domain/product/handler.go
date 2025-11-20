package product

import (
	"strconv"

	"go-learning/internal/database"
	"go-learning/internal/helpers"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service  Service
	response *helpers.Response
}

func NewHandler(service Service, response *helpers.Response) *Handler {
	return &Handler{
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
func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	db := database.GetDB()

	// get body
	var req Product
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
func (h *Handler) ListProduct(c *fiber.Ctx) error {
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
func (h *Handler) GetProduct(c *fiber.Ctx) error {
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
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	db := database.GetDB()

	// get params id
	idStr := c.Params("id")
	id, _ := strconv.Atoi(idStr)

	// get body
	var req Product
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
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
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
