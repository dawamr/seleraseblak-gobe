package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
)

type storeController struct {
	storeService api.StoreService
}

func NewStoreController(service api.StoreService) *storeController {
	return &storeController{
		storeService: service,
	}
}

func (c *storeController) CreateStore(ctx *fiber.Ctx) error {
	store := new(api.Store)
	if err := ctx.BodyParser(store); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.storeService.CreateStore(store); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(store)
}

func (c *storeController) GetStore(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	store, err := c.storeService.GetStore(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Store not found",
		})
	}

	return ctx.JSON(store)
}

func (c *storeController) UpdateStore(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	store := new(api.Store)
	if err := ctx.BodyParser(store); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.storeService.UpdateStore(id, store); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(store)
}

func (c *storeController) DeleteStore(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.storeService.DeleteStore(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *storeController) ListStores(ctx *fiber.Ctx) error {
	stores, err := c.storeService.ListStores(map[string]interface{}{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Filter hanya store yang published
	var activeStores []api.Store
	for _, store := range stores {
		if store.Status == "published" {
			activeStores = append(activeStores, store)
		}
	}

	return ctx.JSON(activeStores)
}
