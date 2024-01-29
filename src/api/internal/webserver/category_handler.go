package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

/**
* NewWebServerCategoryHandler is a function that creates a new WebCategoryHandler struct
* {categoryService} is a pointer to the category service
* returns a category service
**/
func NewWebServerCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

/**
* GetCategories is a function that returns all categories from the database
* {w} is the response writer
* {r} is the request
* returns all categories and an error if something goes wrong
**/
func (wch *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

/**
* GetCategory is a function that returns a category from the database
* {w} is the response writer
* {r} is the request
* returns a category and an error if something goes wrong
**/
func (wch *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	category, err := wch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

/**
* CreateCategory is a function that creates a new category in the database
* {w} is the response writer
* {r} is the request
* returns a category and an error if something goes wrong
**/
func (wch *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := wch.CategoryService.CreateCategory(category.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}