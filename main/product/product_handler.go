package product

import (
	"Products/main/validation"
	"fmt"
	_validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	GetProductById(productId string) (Product, error)
	CreateProduct(product Product) error
}

type ProductHandler struct {
	service   Service
	validator *_validator.Validate
}

func NewProductHandler() ProductHandler {
	_validator.New()
	newProductHandler := ProductHandler{
		service:   ProductService{},
		validator: _validator.New(),
	}

	return newProductHandler
}

func (h *ProductHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/product", h.GetProduct)
	app.Post("/product", h.CreateProduct)
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	book, err := h.service.GetProductById(c.Query("productId"))
	if err != nil {
		return err
	}

	// returns the book with the status 200
	return c.Status(200).JSON(book)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	newProduct := Product{}
	err := c.BodyParser(&newProduct)
	if err != nil {
		return err
	}

	errors := validation.ValidateStruct(newProduct)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// servis vasıtasıyla kitabı create et
	err = h.service.CreateProduct(newProduct)
	if err != nil {
		return err
	}

	fmt.Println(newProduct.Name)
	return c.Status(fiber.StatusCreated).JSON(newProduct)
	return c.SendStatus(201)
}
