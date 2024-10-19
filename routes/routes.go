package routes

import (
	"student-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routes cho sinh viên
	router.GET("/students", controllers.GetStudents)
	router.POST("/students", controllers.CreateStudent)
	router.GET("/students/:id", controllers.GetStudentByID)
	router.PUT("/students/:id", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)

	return router
}
