package product

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,min=3,max=500"`
	Price       int64  `json:"price" validate:"required,min=1"`
	Quantity    int64  `json:"quantity" validate:"required,min=1"`
}
