package database

import (
	"tento/internal/models"
	"tento/internal/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	var dbName string
	isTesting := false
	goEnv := utils.Getenv("GO_ENV", "")
	utils.TentoLogger.Info("GO_ENV: ", goEnv)
	if goEnv == "testing" {
		dbName = "file::memory:?cache=shared"
		//dbName = "testing.db"
		isTesting = true
	} else {
		dbName = utils.Getenv("TENT_DB", "development.db")
	}
	utils.TentoLogger.Info("isTesting?", isTesting)
	var err error
	utils.TentoLogger.Info("Connecting to DB: ", dbName)
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate()
	//if !isTesting {
	populate()
	//}
}

func migrate() {
	utils.TentoLogger.Info("Migrating database ...")
	DB.AutoMigrate(&models.Product{})
}

func populate() {

	var products = []models.Product{{Description: "Nescafe Moka 18g", Stock: 6, Barcode: "7613033458507", BuyPrice: 300, SellPrice: 450},
		{Description: "Vino Gato Cabernet Sauvignon 1/2 Lt Tetra", Stock: 1, Barcode: "7804300122140", BuyPrice: 9031, SellPrice: 11740},
		{Description: "Vino Gato 1000 Cc Blanco", Stock: 1, Barcode: "7804300004019", BuyPrice: 18579, SellPrice: 24150},
		{Description: "Vino Gato 1000 Cc Cabernet Sauvignon", Stock: 1, Barcode: "7804300004002", BuyPrice: 18579, SellPrice: 24150},
		{Description: "Vino Gato 2000 Cc Merlot Tetrax", Stock: 2, Barcode: "7804300131692", BuyPrice: 20785, SellPrice: 27020},
		{Description: "Gato_Dulce-Dc-Vnr750Ccx12-Co", Stock: 0, Barcode: "700808", BuyPrice: 17160, SellPrice: 22310},
		{Description: "Kent Belmont - Rc10", Stock: 2, Barcode: "10070892", BuyPrice: 31454, SellPrice: 40890},
		{Description: "Kent Belmont - Hl20", Stock: 6, Barcode: "10070700", BuyPrice: 31454, SellPrice: 40890},
		{Description: "Kent Neo Ikon Mix Ds90", Stock: 1, Barcode: "10107101", BuyPrice: 27522, SellPrice: 35780},
		{Description: "Pall Mall Azul Hl20S", Stock: 2, Barcode: "10111410", BuyPrice: 25291, SellPrice: 32880},
		{Description: "Pall Mall Boost - Hl10", Stock: 2, Barcode: "10102304", BuyPrice: 31614, SellPrice: 41100},
		{Description: "Pall Mall Boost Xl - Hl20", Stock: 4, Barcode: "10091121", BuyPrice: 27662, SellPrice: 35960},
		{Description: "Pall Clonb 10/200 Kre Sq Chi F&S", Stock: 2, Barcode: "10119670", BuyPrice: 31614, SellPrice: 41100},
		{Description: "Pall Clonb 20/200 Kre Sq Chi F&S", Stock: 5, Barcode: "10119688", BuyPrice: 27662, SellPrice: 35960},
		{Description: "Pall Mall Azul Sc20S", Stock: 12, Barcode: "10111132", BuyPrice: 22130, SellPrice: 28770},
		{Description: "Pall Mall Rojo - Sc20", Stock: 3, Barcode: "10111094", BuyPrice: 22130, SellPrice: 28770},
		{Description: "Escudo Silver Lat 470CC", Stock: 24, Barcode: "7802100002952", BuyPrice: 8474, SellPrice: 10000}}

	var count int64
	DB.Model(&models.Product{}).Count(&count)
	utils.TentoLogger.Info("Found ", count, " products in DB")
	if count < int64(len(products)) {
		utils.TentoLogger.Info("creating Products ...")
		DB.Create(&products)
	}
	//DB.Create(&models.Product{Barcode: "2ASF32415D42", SellPrice: 100})
	//DB.Create(&models.Product{Barcode: "2ASF32415D43", SellPrice: 300})
	//DB.Create(&models.Product{Barcode: "2ASF32415D44", SellPrice: 600})
	//DB.Create(&models.Product{Description: "Nescafe Moka 18g", Stock: 6, Barcode: "7613033458507", BuyPrice: 300, SellPrice: 450})
	//DB.Create(&models.Product{Description: "Vino Gato Cabernet Sauvignon 1/2 Lt Tetra", Stock: 1, Barcode: "7804300122140", BuyPrice: 9031, SellPrice: 11740})
	//DB.Create(&models.Product{Description: "Vino Gato 1000 Cc Blanco", Stock: 1, Barcode: "7804300004019", BuyPrice: 18579, SellPrice: 24150})
	//DB.Create(&models.Product{Description: "Vino Gato 1000 Cc Cabernet Sauvignon", Stock: 1, Barcode: "7804300004002", BuyPrice: 18579, SellPrice: 24150})
	//DB.Create(&models.Product{Description: "Vino Gato 2000 Cc Merlot Tetrax", Stock: 2, Barcode: "7804300131692", BuyPrice: 20785, SellPrice: 27020})
	//DB.Create(&models.Product{Description: "Gato_Dulce-Dc-Vnr750Ccx12-Co", Stock: 0, Barcode: "700808", BuyPrice: 17160, SellPrice: 22310})
	//DB.Create(&models.Product{Description: "Kent Belmont - Rc10", Stock: 2, Barcode: "10070892", BuyPrice: 31454, SellPrice: 40890})
	//DB.Create(&models.Product{Description: "Kent Belmont - Hl20", Stock: 6, Barcode: "10070700", BuyPrice: 31454, SellPrice: 40890})
	//DB.Create(&models.Product{Description: "Kent Neo Ikon Mix Ds90", Stock: 1, Barcode: "10107101", BuyPrice: 27522, SellPrice: 35780})
	//DB.Create(&models.Product{Description: "Pall Mall Azul Hl20S", Stock: 2, Barcode: "10111410", BuyPrice: 25291, SellPrice: 32880})
	//DB.Create(&models.Product{Description: "Pall Mall Boost - Hl10", Stock: 2, Barcode: "10102304", BuyPrice: 31614, SellPrice: 41100})
	//DB.Create(&models.Product{Description: "Pall Mall Boost Xl - Hl20", Stock: 4, Barcode: "10091121", BuyPrice: 27662, SellPrice: 35960})
	//DB.Create(&models.Product{Description: "Pall Clonb 10/200 Kre Sq Chi F&S", Stock: 2, Barcode: "10119670", BuyPrice: 31614, SellPrice: 41100})
	//DB.Create(&models.Product{Description: "Pall Clonb 20/200 Kre Sq Chi F&S", Stock: 5, Barcode: "10119688", BuyPrice: 27662, SellPrice: 35960})
	//DB.Create(&models.Product{Description: "Pall Mall Azul Sc20S", Stock: 12, Barcode: "10111132", BuyPrice: 22130, SellPrice: 28770})
	//DB.Create(&models.Product{Description: "Pall Mall Rojo - Sc20", Stock: 3, Barcode: "10111094", BuyPrice: 22130, SellPrice: 28770})
	//DB.Create(&models.Product{Description: "Escudo Silver Lat 470CC", Stock: 24, Barcode: "7802100002952", BuyPrice: 8474, SellPrice: 10000})
}
