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
