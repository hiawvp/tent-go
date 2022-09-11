package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"tento/internal/database"
	"tento/internal/models"
	"tento/internal/router"
	"tento/internal/utils"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const apiPrefix = "/api/v1"

var r *gin.Engine
var test_prod_id uint
var test_prod models.Product

func init() {
	fmt.Println("Calling init at endpoint_test.go")
	os.Setenv("GO_ENV", "testing")
	utils.SetupLogger()
	database.Setup()
	gin.SetMode(gin.ReleaseMode)
	r = router.Create()
	utils.TentoLogger.SetLogLevel("INFO")
	//database.Setup()
}

func TestHomeEndpoint(t *testing.T) {
	testUrl := apiPrefix + "/"
	//http.Get(testUrl)
	req, _ := http.NewRequest("GET", testUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiVerRedirect(t *testing.T) {
	testUrl := "/"
	//http.Get(testUrl)
	req, _ := http.NewRequest("GET", testUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

func TestGetProducts(t *testing.T) {
	testUrl := apiPrefix + "/products"
	//http.Get(testUrl)
	req, _ := http.NewRequest("GET", testUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	var products []models.Product
	json.Unmarshal(responseData, &products)
	fmt.Println("Found ", len(products), " Products")
	assert.NotEmpty(t, products)
	test_prod_id = products[0].ID
	test_prod = products[0]
	//expectedResponse := "XDDDD"
	//assert.Equal(t, string(responseData), expectedResponse)
	return
}

func TestGetProductById(t *testing.T) {
	if test_prod_id == 0 || test_prod.ID == 0 {
		t.FailNow()
	}
	testUrl := apiPrefix + "/products/" + fmt.Sprintf("%v", test_prod_id)
	//http.Get(testUrl)
	req, _ := http.NewRequest("GET", testUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	var product models.Product
	json.Unmarshal(responseData, &product)
	if assert.NotEmpty(t, product) {
		fmt.Println("Found product :", product.Description)
	}

	assert.Equal(t, product, test_prod)
	return
}

func TestGetProductByBarcode(t *testing.T) {
	if test_prod_id == 0 || test_prod.ID == 0 {
		t.FailNow()
	}
	testUrl := apiPrefix + "/products?barcode=" + test_prod.Barcode
	req, _ := http.NewRequest("GET", testUrl, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	var product models.Product
	json.Unmarshal(responseData, &product)
	if assert.NotEmpty(t, product.ID, "ProdByBarcode should not have ID=0") {
		fmt.Println("Found product :", product.Description)
		assert.Equal(t, test_prod, product)
	}
	return
}

func TestPostProduct(t *testing.T) {
	testUrl := apiPrefix + "/products"
	post_product := models.Product{
		Barcode:     "NANACHI5150",
		BuyPrice:    300,
		SellPrice:   500,
		Stock:       10,
		Description: "Nanachi cool product",
	}
	body, err := json.Marshal(post_product)
	if err != nil {
		fmt.Println("Error! ", err.Error())
		t.FailNow()
	}
	//http.Post(testUrl, "json", bytes.NewBuffer(body))
	req, _ := http.NewRequest("POST", testUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseProduct models.Product
	assert.Equal(t, http.StatusOK, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(responseData, &responseProduct)
	if assert.NotEmpty(t, post_product) {
		fmt.Println("Found post_product :", responseProduct.Description)
	}

	assert.Equal(t, post_product.BuyPrice, responseProduct.BuyPrice)
	assert.Equal(t, post_product.SellPrice, responseProduct.SellPrice)
	assert.NotEqual(t, post_product.ID, responseProduct.ID)
	return
}
