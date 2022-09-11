package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	fmt.Println("Setting TENT_LOG_LEVEL WARN")
	os.Setenv("TENT_LOG_LEVEL", "WARN")
	os.Exit(0)

}
