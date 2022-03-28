package main

import (
	"Products/main/product"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	ProductHandler := product.NewProductHandler()
	ProductHandler.RegisterRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
