package services

import (
	"fmt"
	"tento/internal/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	fmt.Println("Running main from services")
	db = database.DB
}
