package routes

import (
	"go_backend/controllers"

	"github.com/gin-gonic/gin"
)

func FavouriteRoutes(r gin.IRoutes) {
	r.POST("/fav/add", controllers.AddFavourite)
	r.GET("/fav/get/:username", controllers.GetFavourites)
	r.DELETE("/fav/:id", controllers.RemoveFavourite)
}
