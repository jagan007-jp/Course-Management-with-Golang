package routes

import (
	"go_backend/controllers"

	"github.com/gin-gonic/gin"
)

func CourseRoutes(r gin.IRoutes) {

	r.GET("/courses", controllers.FetchCourses)
	r.DELETE("/courses/delete/:id", controllers.DeleteCourse)
	r.POST("/courses/add", controllers.AddCourse)
}
