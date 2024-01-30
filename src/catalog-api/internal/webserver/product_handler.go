package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductService *service.ProductService
}

/**
* NewWebServerProductHandler is a function that creates a new WebProductHandler struct
* {productService} is a pointer to the product service
* returns a product service
**/
func NewWebServerProductHandler(productService *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductService: productService}
}

/**
* GetProducts is a function that returns all products from the database
* {w} is the response writer
* {r} is the request
* returns all products and an error if something goes wrong
**/
func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

/**
* GetProduct is a function that returns a product from the database
* {w} is the response writer
* {r} is the request
* returns a product and an error if something goes wrong
**/
func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	product, err := wph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

/**
* GetProductByCategoryID is a function that returns all products from a category from the database
* {w} is the response writer
* {r} is the request
* returns all products from a category and an error if something goes wrong
**/
func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}

	products, err := wph.ProductService.GetProductByCategoryID(categoryID	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

/**
* CreateProduct is a function that creates a new product in the database
* {w} is the response writer
* {r} is the request
* returns a product and an error if something goes wrong
**/
func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wph.ProductService.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageUrl, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}