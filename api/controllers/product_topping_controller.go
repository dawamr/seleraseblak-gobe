package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
)

type ProductToppingController struct {
    service api.ProductToppingService
}

func NewProductToppingController(service api.ProductToppingService) *ProductToppingController {
    return &ProductToppingController{service: service}
}

// GetProductToppings godoc
// @Summary Get all product toppings
// @Description Get all product toppings with their relations
// @Tags product-toppings
// @Accept json
// @Produce json
// @Success 200 {array} api.ProductTopping
// @Router /product-toppings [get]
func (c *ProductToppingController) GetProductToppings(ctx *fiber.Ctx) error {
    productToppings, err := c.service.GetProductToppings()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(productToppings)
}

// GetProductToppingsByProduct godoc
// @Summary Get product toppings by product ID
// @Description Get all toppings for a specific product
// @Tags product-toppings
// @Accept json
// @Produce json
// @Param productId path int true "Product ID"
// @Success 200 {array} api.ProductTopping
// @Router /products/{productId}/toppings [get]
func (c *ProductToppingController) GetProductToppingsByProduct(ctx *fiber.Ctx) error {
    productID, err := strconv.Atoi(ctx.Params("productId"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
    }
    productToppings, err := c.service.GetProductToppingsByProduct(productID)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(productToppings)
}

// GetProductToppingsByTopping godoc
// @Summary Get product toppings by topping ID
// @Description Get all products for a specific topping
// @Tags product-toppings
// @Accept json
// @Produce json
// @Param toppingId path int true "Topping ID"
// @Success 200 {array} api.ProductTopping
// @Router /toppings/{toppingId}/products [get]
func (c *ProductToppingController) GetProductToppingsByTopping(ctx *fiber.Ctx) error {
    toppingID, err := strconv.Atoi(ctx.Params("toppingId"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid topping ID"})
    }
    productToppings, err := c.service.GetProductToppingsByTopping(toppingID)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(productToppings)
}
