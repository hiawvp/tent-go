package router

import (
	"net/http"
	"tento/internal/handlers"
	"tento/internal/utils"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	var apiPrefix = utils.Getenv("API_URL_PREFIX", "/api/v1/")
	utils.TentoLogger.Info("Api prefix: ", apiPrefix)

	router := gin.Default()
	router.Use(cors.Default())
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, apiPrefix)
	})
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	v1 := router.Group(apiPrefix)
	{
		v1.GET("/", handlers.Home)
		v1.GET("/home", handlers.Home)
		v1.GET("/products/:id", handlers.GetProduct)
		v1.GET("/products", handlers.GetProducts)
		v1.POST("/products", handlers.PostProduct)

	}
	return router
}
