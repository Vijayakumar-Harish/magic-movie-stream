package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"os"
	routes "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	database "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	"github.com/gin-contrib/cors"
)

func main() {
	
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
			log.Println("Allowed Origin:", origins[i])
		}
	} else {
		origins = []string{"http://localhost:5173"}
		log.Println("Allowed Origin: http://localhost:5173")
	}

	config := cors.Config{}
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET","POST","PATCH","PUT","DELETE","OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type","Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	router.Use(gin.Logger())

	var client *mongo.Client = database.Connect()

	routes.SetupUnProtectedRoutes(router, client)
	routes.SetupProtectedRoutes(router, client)
	

	if err:=router.Run(":8082");err!=nil{
		fmt.Println("Failed to start server", err)
	}
}