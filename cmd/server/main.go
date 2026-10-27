package main

import (
	"log"

	_ "github.com/carlosfdsn/sacoli/docs"
	"github.com/carlosfdsn/sacoli/internal/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title sacoli API
// @version 1.0
// @description  API para gerenciamento de compras.
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", handlers.HealthCheck)

	itemsGroup := r.Group("/items")
	{
		itemsGroup.GET("", handlers.ListItems)
		itemsGroup.GET("/:id", handlers.GetItemByID)
		itemsGroup.POST("", handlers.CreateItem)
		itemsGroup.PUT("/:id", handlers.UpdateItem)
		itemsGroup.DELETE("/:id", handlers.DeleteItem)
		itemsGroup.PATCH("/:id/bought", handlers.MarkAsBought)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}