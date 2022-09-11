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

func FindProducts(page, limit int, search string) ([]models.Product, error) {
	utils.TentoLogger.Info("FindProducts")
	// TODO: paginate resaults
	var products []models.Product
	result := database.DB.
		Where("description LIKE ?", fmt.Sprintf("%%%v%%", search)).
		Limit(limit).
		Offset(limit * page).
		Find(&products)
	if result.Error != nil {
		utils.TentoLogger.Error("Error!! ", result.Error.Error())
	}
	utils.TentoLogger.Info(fmt.Sprintf("Found %v products", len(products)))
	return products, result.Error

}

func FindProductById(id int64) (models.Product, error) {
	utils.TentoLogger.Info("FindProductById")
	// TODO: paginate resaults
	var product models.Product
	result := database.DB.First(&product, id)
	if result.Error != nil {
		utils.TentoLogger.Error("Error!! ", result.Error.Error())
	}
	utils.TentoLogger.Info("Found product", product.Barcode)
	return product, result.Error
}

func FindProductByBarcode(barcode string) (models.Product, error) {
	utils.TentoLogger.Info("FindProductByBarcode")
	var product models.Product
	result := database.DB.Where("barcode = ?", barcode).First(&product)
	if result.Error != nil {
		utils.TentoLogger.Error("Error!! ", result.Error.Error())
	} else {
		utils.TentoLogger.Info("Found product", product.Barcode)
	}
	return product, result.Error

}
