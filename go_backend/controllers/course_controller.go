package controllers

import (
	"context"
	"go_backend/db"
	"go_backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchCourses(c *gin.Context) {
	limit1 := c.DefaultQuery("limit", "6")
	offset1 := c.DefaultQuery("offset", "0")

	limit, err1 := strconv.Atoi(limit1)
	offset, err2 := strconv.Atoi(offset1)

	if err1 != nil || err2 != nil || limit < 0 || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit or offset"})
		return
	}

	var courses []models.Course
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("courses")

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))

	cursor, err := collection.Find(ctx, bson.M{}, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cursor.All(ctx, &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	role := c.MustGet("role").(string)
	if role != "admin" {
		c.JSON(http.StatusOK, gin.H{"courses": courses})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses, "role": role})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection("courses")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, bson.M{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bson.M{"message": "Course deleted successfully"})
}

func AddCourse(c *gin.Context) {
	var course_data models.Course
	if err := c.ShouldBindJSON(&course_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	collection := db.DB.Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, course_data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course added the courses successfully"})
}
