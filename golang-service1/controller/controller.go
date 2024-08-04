package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang-service1/model"
	"golang-service1/service"
	"net/http"
	"strconv"
)

func CreateCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := service.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	category, err := service.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Gọi API lấy danh sách sản phẩm từ ProductService
	products, err := getProductsByCategoryId(category.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	category.Products = products

	c.JSON(http.StatusOK, category)
}

func getProductsByCategoryId(categoryId int64) ([]model.Product, error) {
	resp, err := http.Get("http://localhost:8080/products/category/" + strconv.FormatInt(categoryId, 10))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var products []model.Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}
