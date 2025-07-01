package main

import (
	"go_backend/db"
	"go_backend/middleware"
	"go_backend/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.UserRoutes(r)
	protected := r.Group("/")
	protected.Use(middleware.ValidateJWT())
	routes.CourseRoutes(protected)
	routes.FavouriteRoutes(protected)
	r.Run(":8080")
}
