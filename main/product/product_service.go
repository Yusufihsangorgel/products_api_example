package product

type ProductService struct {
	// db (repository)
}

func (bs ProductService) CreateProduct(product Product) error {
	return nil
}

func (bs ProductService) GetProductById(productId string) (Product, error) {

	// this book will be taken from the db
	product := Product{
		Id:          "",
		Name:        "Book",
		Description: "Book Writer",
		Price:       100,
		Quantity:    10,
	}

	product.Id = "1"

	return product, nil
}
