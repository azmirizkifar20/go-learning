package controllers

import (
	"go-learning/internal/helpers"
	"go-learning/internal/services"

	"github.com/gofiber/fiber/v2"
)

type StoreController struct {
	storeService *services.StoreService
	response     *helpers.Response
}

func NewStoreController(storeService *services.StoreService, response *helpers.Response) *StoreController {
	return &StoreController{
		storeService: storeService,
		response:     response,
	}
}

/*
===========================================================================================
Method: GET
URL Path: /v1/stores
===========================================================================================
*/
func (h *StoreController) GetStoreNoo(c *fiber.Ctx) error {
	// Implementasi handler untuk mendapatkan data Store Noo
	return h.response.Send(c, fiber.StatusOK, nil, "GetStoreNoo not implemented", nil)
}
