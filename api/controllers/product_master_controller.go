package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seleraseblak/backend/api"
)

type productMasterController struct {
	productMasterService api.ProductMasterService
}

func NewProductMasterController(service api.ProductMasterService) *productMasterController {
	return &productMasterController{
		productMasterService: service,
	}
}

func (c *productMasterController) CreateProductMaster(ctx *fiber.Ctx) error {
	product := new(api.ProductMaster)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.productMasterService.CreateProductMaster(product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (c *productMasterController) GetProductMaster(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product, err := c.productMasterService.GetProductMaster(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product master not found",
		})
	}

	return ctx.JSON(product)
}

func (c *productMasterController) UpdateProductMaster(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	product := new(api.ProductMaster)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.productMasterService.UpdateProductMaster(id, product); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(product)
}

func (c *productMasterController) DeleteProductMaster(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.productMasterService.DeleteProductMaster(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *productMasterController) ListProductMasters(ctx *fiber.Ctx) error {
	params := make(map[string]interface{})

	// Add pagination
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	params["page"] = page
	params["limit"] = limit

	// Add filters
	if search := ctx.Query("search"); search != "" {
		params["search"] = search
	}

	// Add category filter
	if category := ctx.Query("category"); category != "" {
		params["category"] = category
	}

	products, err := c.productMasterService.ListProductMasters(params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data":  products,
		"page":  page,
		"limit": limit,
	})
}
