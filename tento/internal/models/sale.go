package models

import (
	"time"

	"gorm.io/gorm"
)

// Sale belongs to an User
// has many Product
// could have a Payment field
type Sale struct {
	gorm.Model
	SubTotal int
	Total    int
	Date     time.Time
	Items    []SaleProduct `gorm:"many2many:sale_products;"`
	UserID   int
	User     User
}
