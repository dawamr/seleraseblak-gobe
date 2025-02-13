package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
)

type SpicyLevelController struct {
	spicyLevelService api.SpicyLevelService
}

func NewSpicyLevelController(service api.SpicyLevelService) *SpicyLevelController {
	return &SpicyLevelController{spicyLevelService: service}
}

func (c *SpicyLevelController) GetSpicyLevels(ctx *fiber.Ctx) error {
	spicyLevels, err := c.spicyLevelService.GetSpicyLevels()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(spicyLevels)
}

func (c *SpicyLevelController) GetSpicyLevel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	spicyLevel, err := c.spicyLevelService.GetSpicyLevel(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Spicy level not found",
		})
	}
	return ctx.JSON(spicyLevel)
}
