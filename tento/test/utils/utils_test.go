package utils

import (
	"tento/pkg/models"
	"tento/pkg/utils"
	"testing"
)

func TestPrettyPrintEmpty(t *testing.T) {
	var i interface{}
	got := utils.PrettyPrint(i)
	want := "null"
	if got != want {
		t.Error("TestPrettyPrintEmpty error! Got: ", got, ", Wanted: ", want)
	}
}

func TestPrettyPrintProduct(t *testing.T) {
	prod := models.Product{Barcode: "2ASF32415D44", SellPrice: 600}
	got := utils.PrettyPrint(prod)
	//fmt.Println(got)
	if got == "null" || len(got) == 0 {
		t.Error("TestPrettyPrintProduct error! Got: ", got)
	}
}
