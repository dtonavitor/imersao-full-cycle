package database

import (
	"database/sql"

	"github.com/dtonavitor/imersao-full-cycle/src/api/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

/**
* NewCategoryDB is a function that creates a new CategoryDB struct
* {db} is a pointer to the database connection
* returns a database connection for categories
**/
func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

/**
* GetCategories is a function that returns all categories from the database
* returns all categories and an error if something goes wrong
**/
func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err:= rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	return categories, nil
}

/**
* GetCategory is a function that returns a category from the database
* {id} is the id of the category
* returns a category and an error if something goes wrong
**/
func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

/**
* GetCategory is a function that returns a category from the database
* {id} is the id of the category
* returns a category and an error if something goes wrong
**/
func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories (id, name) VALUES (?, ?)", category.ID, category.Name)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}