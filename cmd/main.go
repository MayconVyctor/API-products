package main

import (
	"go-api/controler"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//Camada de controls
	ProductControler := controler.NewProductControler(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductControler.GetProducts)
	server.POST("/product", ProductControler.CreateProduct)
	server.PUT("/product/:productId", ProductControler.UpdateProduct)
	server.GET("/product/:productId", ProductControler.GetProductById)
	server.DELETE("/product/:productId", ProductControler.DeleteProduct)

	server.Run(":8000")
}
