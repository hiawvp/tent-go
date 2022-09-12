package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Barcode     string `gorm:"unique"`
	BuyPrice    uint
	SellPrice   uint
	Stock       uint
	Category    string
	Description string
}

// `json:"col_name"`  ?

// belongs to a Product
type SaleProduct struct {
	gorm.Model
	Quantity  uint
	BoughtFor uint
	SoldFor   uint
	ProductID int
	Product   Product
}
