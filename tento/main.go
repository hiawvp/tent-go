package main

import (
	"log"

	"tento/internal/database"
	"tento/internal/router"
	"tento/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	utils.SetupLogger()
	utils.TentoLogger.Info("Running in '", utils.Getenv("GO_ENV", "development"), "' mode.")

	database.Setup()

	ginMode := utils.Getenv("GIN_MODE", gin.DebugMode)
	gin.SetMode(ginMode)
	router := router.Create()
	err := router.Run()
	if err != nil {
		log.Fatal("Could not launch router!\n", err.Error())
	}

}
