package entity

import "github.com/google/uuid"

type Category struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

/**
* NewCategory is a function that creates a new Category struct
* {name} is the name of the category
* returns a new category
**/
func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CategoryID string `json:"category_id"`
	ImageUrl string `json:"image_url"`
}

/**
* NewProduct is a function that creates a new Product struct
* {name} is the name of the product
* {description} is the description of the product
* {price} is the price of the product
* {categoryID} is the id of the category of the product
* {imageUrl} is the url of the image of the product
* returns a new product
**/
func NewProduct(name, description, categoryID, imageUrl string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
		ImageUrl:    imageUrl,
	}
}