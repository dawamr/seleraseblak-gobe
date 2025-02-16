package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/seleraseblak/backend/api/controllers"
	"github.com/seleraseblak/backend/config"
	"github.com/seleraseblak/backend/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Initialize services
	storeService := services.NewStoreService(db)
	productService := services.NewProductService(db)
	productMasterService := services.NewProductMasterService(db)
	toppingService := services.NewToppingService(db)
	spicyLevelService := services.NewSpicyLevelService()
	productToppingService := services.NewProductToppingService(db)

	// Initialize controllers
	storeController := controllers.NewStoreController(storeService)
	productController := controllers.NewProductController(productService)
	productMasterController := controllers.NewProductMasterController(productMasterService)
	toppingController := controllers.NewToppingController(toppingService)
	spicyLevelController := controllers.NewSpicyLevelController(spicyLevelService)
	productToppingController := controllers.NewProductToppingController(productToppingService)

	// Create Fiber app
	app := fiber.New()

	// CORS middleware dengan konfigurasi yang lebih lengkap
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,http://localhost:4321,https://seleraseblak-website.pages.dev,https://seleraseblak.com", // URL frontend yang diizinkan
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,HEAD,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
		ExposeHeaders: "Content-Length, Access-Control-Allow-Origin",
		MaxAge: 86400, // 24 jam dalam detik
	}))

	// API routes
	api := app.Group("/api")

	// Store routes
	stores := api.Group("/stores")
	stores.Post("/", storeController.CreateStore)
	stores.Get("/:id", storeController.GetStore)
	stores.Put("/:id", storeController.UpdateStore)
	stores.Delete("/:id", storeController.DeleteStore)
	stores.Get("/", storeController.ListStores)

	// Product routes
	stores.Post("/:store_id/products", productController.CreateProduct)
	stores.Get("/:store_id/products/:id", productController.GetProduct)
	stores.Put("/:store_id/products/:id", productController.UpdateProduct)
	stores.Delete("/:store_id/products/:id", productController.DeleteProduct)
	stores.Get("/:store_id/products", productController.ListProducts)

	// Product Master routes
	productMasters := api.Group("/product-masters")
	productMasters.Post("/", productMasterController.CreateProductMaster)
	productMasters.Get("/:id", productMasterController.GetProductMaster)
	productMasters.Put("/:id", productMasterController.UpdateProductMaster)
	productMasters.Delete("/:id", productMasterController.DeleteProductMaster)
	productMasters.Get("/", productMasterController.ListProductMasters)

	// Topping routes
	api.Get("/toppings", toppingController.GetToppings)
	api.Get("/toppings/:id", toppingController.GetTopping)

	// Spicy Level routes
	api.Get("/spicy-levels", spicyLevelController.GetSpicyLevels)
	api.Get("/spicy-levels/:id", spicyLevelController.GetSpicyLevel)

	// Product Topping routes
	api.Get("/product-toppings", productToppingController.GetProductToppings)
	api.Get("/products/:productId/toppings", productToppingController.GetProductToppingsByProduct)
	api.Get("/toppings/:toppingId/products", productToppingController.GetProductToppingsByTopping)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
