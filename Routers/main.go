package Routers

import (
	ScheduleCourseController "Project1/Controller/ScheduleCourse"
	"github.com/gin-gonic/gin"
)

func ScheduleCourseRouter(r *gin.Engine) {
	g := r.Group("/api/v1")
	g.POST("/course/schedule", ScheduleCourseController.ScheCourseController{}.ScheduleCourse)
}
