package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"tento/internal/models"
	"tento/internal/services"
	"tento/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// TODO: paginate resaults
func GetProducts(c *gin.Context) {
	utils.TentoLogger.Info("called")
	if barcode := c.Query("barcode"); len(barcode) > 0 {
		getProductByBarcode(c, barcode)
	} else {
		getFilteredProducts(c)
	}
}

func getProductByBarcode(c *gin.Context, barcode string) {
	utils.TentoLogger.Info("barcode: ", barcode)
	product, err := services.FindProductByBarcode(barcode)
	if err != nil {
		utils.TentoLogger.Error("barcode error: ", err.Error())
		//errResponse := NewCustomErrorResponse(err, "barcode", barcode)
		//c.JSON(errResponse.httpStatusCode, errResponse.body)
		c.JSON(http.StatusNotFound, gin.H{"code": "ITEM_NOT_FOUND", "message": "xddd"})
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func getFilteredProducts(c *gin.Context) {
	limit := utils.ParseIntDefault(c.Query("limit"), 10)
	page := utils.ParseIntDefault(c.Query("page"), 0)
	search := c.Query("search")
	utils.TentoLogger.Info("Queryargs: ", limit, page, search)
	products, _ := services.FindProducts(page, limit, search)
	if len(products) == 0 {
		utils.TentoLogger.Warn("found 0 products")
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	//utils.TentoLogger.Info("GetProduct by id")
	//utils.TentoLogger.Error("GetProduct by id")

	//TODO: extract url param validation to func
	str_id := c.Param("id")
	id, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		c.JSON(
			400,
			gin.H{
				"code":    "BAD_REQUEST",
				"message": fmt.Sprintf("Invalid ID : %v", str_id)})
		return
	}

	product, err := services.FindProductById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(
			404,
			gin.H{
				"code":    "ITEM_NOT_FOUND",
				"message": fmt.Sprintf("No Product with ID : %v", str_id)})
		return
	}
	c.JSON(http.StatusOK, product)
}

func PostProduct(c *gin.Context) {
	utils.TentoLogger.Info("PostProduct ")

	var product models.Product
	// TODO: fields validator
	err := c.ShouldBindJSON(&product)
	if err != nil {
		utils.TentoLogger.Error("Error binding body json ", err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	prod, err := services.CreateProduct(product)
	if err != nil {
		if sqlite3Err := err.(sqlite3.Error); errors.Is(err, sqlite3Err) {
			fmt.Println("Kek: ", sqlite3Err.Code)
		}
		utils.TentoLogger.Error("Could not create product. Error: ", err.Error())
		c.JSON(
			404,
			gin.H{
				"code":    "COULD_NOT_CREATE",
				"message": err.Error()})
		return
	}
	utils.TentoLogger.Info("Crated Product ", prod)
	c.JSON(http.StatusOK, prod)
}
