package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"tento/internal/models"
	"tento/internal/services"
	"tento/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	c.String(http.StatusOK, "XD")
}

func GetSale(c *gin.Context) {
	str_id := c.Param("id")
	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		errHint := fmt.Sprintf("Invalid ID : '%v'. (%v)", str_id, err.Error())
		abortRequest(c, http.StatusBadRequest, "BAD_REQUEST", errHint)
		return
	}
	getSaleById(c, id)

}

func getSaleById(c *gin.Context, saleId int64) {
	utils.TentoLogger.Info("saleId: ", saleId)
	sale, err := services.FindSaleById(saleId)
	if err != nil {
		errHint := fmt.Sprintf("No Sale with ID : '%v'. (%v)", saleId, err.Error())
		abortRequest(c, http.StatusNotFound, "NOT_FOUND", errHint)
		return
	}
	c.JSON(http.StatusOK, sale)
}

// start a new sale
func PostSale(c *gin.Context) {
	utils.TentoLogger.Info("")
	// who is selling? get user by cookie / url / json
	sale := models.Sale{}

	if barcode := c.Query("barcode"); barcode != "" {
		qty := utils.ParseIntDefault(c.Query("qty"), 1)
		if prod, err := services.FindProductByBarcode(barcode); err == nil {
			// update stock
			saleProd := models.SaleProduct{
				Product:   prod,
				Quantity:  uint(qty),
				BoughtFor: prod.BuyPrice,
				SoldFor:   prod.SellPrice,
			}
			utils.TentoLogger.Info("SaleProd: ", saleProd)
			//
			sale.Items = append(sale.Items, saleProd)
		}
	}
	utils.TentoLogger.Info("Sale: ", sale)
	//err := c.ShouldBindJSON(&sale)
	//if err != nil {
	//utils.TentoLogger.Error("Error binding body json ", err.Error())
	//abortRequest(c, http.StatusBadRequest, "BINDING_ERROR", err.Error())
	//return
	//}

	//err = services.CreateSale(&sale)
	//if err != nil {
	////if sqlite3Err := err.(sqlite3.Error); errors.Is(err, sqlite3Err) {
	////utils.TentoLogger.Error("sqlite3Err: ", sqlite3Err.Code)
	////}
	//utils.TentoLogger.Error("Could not create sale. Error: ", err.Error())
	//abortRequest(c, http.StatusNotFound, "POST_ERROR", err.Error())
	//return
	//}
	utils.TentoLogger.Info("Created sale ", sale)
	c.JSON(http.StatusOK, sale)
}

//func PatchSale(c *gin.Context) {
//str_id := c.Param("id")
//id, err := strconv.ParseInt(str_id, 10, 64)
//if err != nil {
//errHint := fmt.Sprintf("Invalid ID : '%v'. (%v)", str_id, err.Error())
//abortRequest(c, http.StatusBadRequest, "BAD_REQUEST", errHint)
//return
//}
//utils.TentoLogger.Info("")
//var sale models.Sale
//// TODO: fields validator
//err := c.ShouldBindJSON(&sale)
//if err != nil {
//utils.TentoLogger.Error("Error binding body json ", err.Error())
//abortRequest(c, http.StatusBadRequest, "BINDING_ERROR", err.Error())
//return
//}

//err = services.CreateSale(&sale)
//if err != nil {
////if sqlite3Err := err.(sqlite3.Error); errors.Is(err, sqlite3Err) {
////utils.TentoLogger.Error("sqlite3Err: ", sqlite3Err.Code)
////}
//utils.TentoLogger.Error("Could not create sale. Error: ", err.Error())
//abortRequest(c, http.StatusNotFound, "POST_ERROR", err.Error())
//return
//}
//utils.TentoLogger.Info("Created sale ", sale)
//c.JSON(http.StatusOK, sale)
//}
