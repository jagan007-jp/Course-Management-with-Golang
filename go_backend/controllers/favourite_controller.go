package controllers

import (
	"context"
	"go_backend/db"
	"go_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func AddFavourite(c *gin.Context) {
	var fav models.Favourites
	if err := c.ShouldBindJSON(&fav); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	collection := db.DB.Collection("favourites")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": fav.ID, "username": fav.Username}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "Already added to favourites"})
		return
	}

	_, err = collection.InsertOne(ctx, fav)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add favourite"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Added to favourites"})
}

func GetFavourites(c *gin.Context) {
	username := c.Param("username")
	var favs []models.Favourites

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("favourites")
	cursor, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	if err := cursor.All(ctx, &favs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Successfully retrieved favourites",
		"favCourses": favs,
	})
}

func RemoveFavourite(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("favourites")
	_, err := collection.DeleteOne(ctx, bson.M{"id": id, "username": body.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course removed from favourites"})
}
