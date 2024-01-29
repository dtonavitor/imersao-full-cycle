package database

import (
	"database/sql"

	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

/**
* NewProductDB is a function that creates a new ProductDB struct
* {db} is a pointer to the database connection
* returns a database connection for products
**/
func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

/**
* GetProducts is a function that returns all products from the database
* returns all products and an error if something goes wrong
**/
func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

/**
* GetProduct is a function that returns a product from the database
* {id} is the id of the product
* returns a product and an error if something goes wrong
**/
func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := pd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageUrl)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

/**
* GetProductByCategoryID is a function that returns all products from a category from the database
* {categoryID} is the id of the category
* returns all products from a category and an error if something goes wrong
**/
func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, category_id, image_url FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageUrl); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

/**
* CreateProduct is a function that creates a product in the database
* {product} is the product to be created
* returns an error if something goes wrong
**/
func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.CategoryID, product.ImageUrl)
	if err != nil {
		return nil, err
	}
	return product, nil
}