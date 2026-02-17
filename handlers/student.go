package handlers

import (
	"go-study/database" // 프로젝트 경로에 맞춰 수정
	"go-study/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students) // 전체 조회
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": students})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}
	database.DB.Create(&student) // DB 저장
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": student})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// 1. 해당 ID가 있는지 확인
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "학생 없음"})
		return
	}

	// 2. 입력 데이터 바인딩
	var input models.UpdateStudentInput
	c.ShouldBindJSON(&input)

	// 3. 업데이트 (Updates는 맵이나 구조체로 변경된 필드만 반영함)
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
