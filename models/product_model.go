package models

type product struct {
	ID    int 		`json:"id"`
	Name  string 	`json:"name" validate:"required,min=3,max=20"`
	Price float64 	`json:"price" validate:"required,min=1"`
}





