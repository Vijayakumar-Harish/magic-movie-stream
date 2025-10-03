package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	routes "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
)

func main() {
	
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	routes.SetupUnProtectedRoutes(router)
	routes.SetupProtectedRoutes(router)
	

	if err:=router.Run(":8082");err!=nil{
		fmt.Println("Failed to start server", err)
	}
}