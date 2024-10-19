package controllers

import (
	"net/http"
	"student-management/config"
	"student-management/models"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{"data": students})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func GetStudentByID(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&student)
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found!"})
		return
	}

	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"data": "Student deleted successfully"})
}
