package controllers

import (
	"context"
	"net/http"
	"time"

	database "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/database"
	model "github.com/Vijayakumar-Harish/MagicStreamMovies/Server/MagicStreamMoviesServer/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var movieCollection *mongo.Collection = database.OpenCollection("movies")
func GetMovies() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.JSON(200, gin.H{"message":"List of movies"}) // gin.H is for return JSON structure

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var movies [] model.Movie

		cursor, err := movieCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies."})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode movies."})
		}

		c.JSON(http.StatusOK, movies)
	}
}