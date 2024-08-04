package repository

import (
	"golang-service1/db"
	"golang-service1/model"
)

func CreateCategory(category model.Category) (int64, error) {
	var id int64
	err := db.DB.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id", category.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetCategory(id int64) (model.Category, error) {
	var category model.Category
	err := db.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}
