package services

import (
	"errors"
	"fmt"
	"tento/internal/database"
	"tento/internal/models"
	"tento/internal/utils"

	"github.com/mattn/go-sqlite3"
)

func CreateProduct(prod models.Product) (models.Product, error) {
	utils.TentoLogger.Info("CreateProduct")

	result := database.DB.Create(&prod)
	if result.Error != nil {
		if sqlite3Err := result.Error.(sqlite3.Error); errors.Is(result.Error, sqlite3Err) {
			utils.TentoLogger.Error("Sqlite3 error! ", sqlite3Err.Error())
		}
		utils.TentoLogger.Error("Error creating product ", result.Error.Error())
	}
	return prod, result.Error
}

func FindProducts() []models.Product {
	// TODO: paginate resaults
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		utils.TentoLogger.Error("Error!! ", result.Error.Error())
	}
	utils.TentoLogger.Info(fmt.Sprintf("Found %v products", len(products)))
	return products

}

func FindProductById(id int64) (models.Product, error) {
	// TODO: paginate resaults
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		utils.TentoLogger.Error("Error!! ", result.Error.Error())
	}
	utils.TentoLogger.Info(fmt.Sprintf("Found product", product.Barcode))
	return product, result.Error

}
