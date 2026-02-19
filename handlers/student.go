package handlers

import (
	"go-study/database"
	"go-study/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": students})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": student})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "학생 없음"})
		return
	}

	var input models.UpdateStudentInput
	c.ShouldBindJSON(&input)

	database.DB.Model(&student).Updates(input)
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": student})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "삭제 실패"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "삭제 완료"})
}

func UploadProfile(c *gin.Context) {
	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "파일을 찾을 수 없습니다."})
		return
	}

	dst := "uploads/" + id + "_" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "파일 저장 실패"})
		return
	}

	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "학생을 찾을 수 없습니다."})
		return
	}

	database.DB.Model(&student).Update("ProfileImage", dst)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"path":   dst,
	})
}
