package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken 	   string
}

var users = map[string]Login{}

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if(err != nil) {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada usecase
	ProductUsecase := usecase.NewProcuctUsecase(ProductRepository)

	//Camada de Controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	// server.POST("/register", register)
	// server.POST("/login", login)
	// server.POST("/logout", logout)
	// server.POST("/protected", protected)

	server.Run(":8000")
}

func register(r *http.Request) {

}