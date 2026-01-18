package main

import (
	controller "go-products/contoller"
	"go-products/db"
	"go-products/repository"
	"go-products/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectionDB()

	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductById)
	server.POST("/products", ProductController.AddProduct)
	server.DELETE("/products/:id", ProductController.DeleteProductById)
	server.Run(":8080")
}
