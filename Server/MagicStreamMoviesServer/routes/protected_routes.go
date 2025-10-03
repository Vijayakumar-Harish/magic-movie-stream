package routes

import (
	controller "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/controllers"
	middleware "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(router *gin.Engine){
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.GET("/recommendedmovies", controller.GetRecommendedMovies())
}