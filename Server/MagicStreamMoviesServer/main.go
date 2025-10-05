package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	routes "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"

	database "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/database"

)

func main() {
	
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	var client *mongo.Client = database.Connect()

	routes.SetupUnProtectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)
	

	if err:=router.Run(":8082");err!=nil{
		fmt.Println("Failed to start server", err)
	}
}