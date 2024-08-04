package service

import (
	"golang-service1/model"
	"golang-service1/repository"
)

func CreateCategory(category model.Category) (int64, error) {
	return repository.CreateCategory(category)
}

func GetCategory(id int64) (model.Category, error) {
	return repository.GetCategory(id)
}
