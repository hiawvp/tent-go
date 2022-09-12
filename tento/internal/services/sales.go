package services

import (
	"errors"
	"tento/internal/database"
	"tento/internal/models"
	"tento/internal/utils"

	"github.com/mattn/go-sqlite3"
)

func FindSaleById(id int64) (models.Sale, error) {
	utils.TentoLogger.Info(" id: ", id)
	var sale models.Sale
	result := database.DB.First(&sale, id)
	if result.Error != nil {
		utils.TentoLogger.Error(result.Error.Error())
	}
	utils.TentoLogger.Info("Found sale with total", sale.Total)
	return sale, result.Error
}

func CreateSale(sale *models.Sale) error {
	utils.TentoLogger.Info("")

	result := database.DB.Create(&sale)
	if result.Error != nil {
		if sqlite3Err := result.Error.(sqlite3.Error); errors.Is(result.Error, sqlite3Err) {
			utils.TentoLogger.Error("Sqlite3 error! ", sqlite3Err.Error())
		}
		utils.TentoLogger.Error("Error creating sale ", result.Error.Error())
	}
	return result.Error
}
