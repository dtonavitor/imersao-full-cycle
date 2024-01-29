package service

import (
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/database"
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

/**
* NewProductService is a function that creates a new ProductService struct
* {productDB} is a pointer to the database connection
* returns a database connection for products
**/
func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

/**
* GetProducts is a function that returns all products from the database
* returns all products and an error if something goes wrong
**/
func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

/**
* GetProduct is a function that returns a product from the database
* {id} is the id of the product
* returns a product and an error if something goes wrong
**/
func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

/**
* GetProductByCategoryID is a function that returns all products from a category from the database
* {categoryID} is the id of the category
* returns all products from a category and an error if something goes wrong
**/
func (ps *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

/**
* CreateProduct is a function that creates a new product in the database
* {name} is the name of the product
* {description} is the description of the product
* {category_id} is the id of the category of the product
* {image_url} is the url of the image of the product
* {price} is the price of the product
* returns the created product and an error if something goes wrong
**/
func (ps *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}