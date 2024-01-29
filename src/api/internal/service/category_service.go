package service

import (
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/database"
	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

/**
* NewCategoryService is a function that creates a new CategoryService struct
* {categoryDB} is a pointer to the database connection
* returns a database connection for categories
**/
func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

/**
* GetCategories is a function that returns all categories from the database
* returns all categories and an error if something goes wrong
**/
func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

/**
* CreateCategory is a function that creates a new category in the database
* {name} is the name of the category
* returns the created category and an error if something goes wrong
**/
func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := cs.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

/**
* GetCategory is a function that returns a category from the database
* {id} is the id of the category
* returns a category and an error if something goes wrong
**/
func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}