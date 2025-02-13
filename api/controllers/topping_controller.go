package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
)

type ToppingController struct {
    toppingService api.ToppingService
}

func NewToppingController(ts api.ToppingService) *ToppingController {
    return &ToppingController{toppingService: ts}
}

func (tc *ToppingController) GetToppings(c *fiber.Ctx) error {
    toppings, err := tc.toppingService.GetToppings()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(toppings)
}

func (tc *ToppingController) GetTopping(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid ID format",
        })
    }

    topping, err := tc.toppingService.GetTopping(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Topping not found",
        })
    }

    return c.JSON(topping)
}
