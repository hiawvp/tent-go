package database

import (
	"errors"
	"fmt"
	"os"
	"tento/internal/database"
	"tento/internal/models"
	"tento/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var test_db *gorm.DB
var test_prod = models.Product{Barcode: "1234F", SellPrice: 500}
var test_prod_id uint

func init() {
	os.Setenv("GO_ENV", "testing")
	utils.SetupLogger()
	database.Setup()
	test_db = database.DB
}

func TestCreateItem(t *testing.T) {
	test_db.Create(&test_prod)
	product := models.Product{}
	test_db.First(&product, "barcode = ?", test_prod.Barcode) // find product with code D42
	assert.Equal(t, product.SellPrice, test_prod.SellPrice)
	assert.NotEqual(t, product.CreatedAt, test_prod.CreatedAt)
	test_prod_id = product.ID
}

func TestItemByIDNotFound(t *testing.T) {
	product := models.Product{}
	result := test_db.First(&product, 999)
	assert.True(t, errors.Is(result.Error, gorm.ErrRecordNotFound))
}

func TestUpdateItem(t *testing.T) {
	if test_prod_id == 0 {
		t.FailNow()
	}

	product := models.Product{}
	test_db.First(&product, test_prod_id)
	assert.Equal(t, product.ID, test_prod_id)

	result := test_db.Model(&product).Update("SellPrice", 200)
	assert.Equal(t, result.RowsAffected, int64(1))
}

func TestDeleteItem(t *testing.T) {
	if test_prod_id == 0 {
		t.FailNow()
	}

	product := models.Product{}
	result := test_db.Delete(&product, 1)
	assert.Equal(t, result.RowsAffected, int64(1))
}

func TestDeleteAllProducts(t *testing.T) {
	// permanent delete
	result := test_db.Unscoped().Where("1 = 1").Delete(&models.Product{})
	// soft delete:
	//result := test_db.Where("1 = 1").Delete(&models.Product{})
	fmt.Println("Affected ", result.RowsAffected, " Rows")
	assert.Greater(t, result.RowsAffected, int64(0))
}

func TestEmptyDb(t *testing.T) {
	var products []models.Product
	test_db.Find(&products)
	assert.Empty(t, products)
}
