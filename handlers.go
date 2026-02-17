package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student 구조체: 생성(POST) 시에는 모든 값이 필수입니다.
type Student struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
	Dept string `json:"dept"`
}

// 🆕 UpdateStudentInput: 수정(PATCH) 시에는 필드들이 선택사항입니다.
// required를 제거하여 필요한 데이터만 보낼 수 있게 합니다.
type UpdateStudentInput struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"omitempty,gt=0"` // 값이 있을 때만 0보다 큰지 검사
	Dept string `json:"dept"`
}

var students = []Student{
	{ID: 1, Name: "Kim Junseung", Age: 24, Dept: "Smart Software"},
	{ID: 2, Name: "Min-ji", Age: 22, Dept: "Computer Science"},
	{ID: 3, Name: "Seung-woo", Age: 25, Dept: "AI Engineering"},
	{ID: 4, Name: "Ji-hyeon", Age: 23, Dept: "Smart Software"},
	{ID: 5, Name: "Do-yun", Age: 21, Dept: "Electronic Engineering"},
	{ID: 6, Name: "Ha-rin", Age: 22, Dept: "Nursing"},
	{ID: 7, Name: "Jun-ho", Age: 26, Dept: "Mechanical Engineering"},
	{ID: 8, Name: "Ye-rin", Age: 20, Dept: "Design"},
}

func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": students})
}

func CreateStudent(c *gin.Context) {
	var newStudent Student
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	for _, s := range students {
		if s.ID == newStudent.ID {
			c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "중복된 ID"})
			return
		}
	}

	students = append(students, newStudent)
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": newStudent})
}

// [PATCH] 수정 로직 보강
func UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")

	// Student 대신 UpdateStudentInput 사용 (Validation 에러 해결)
	var input UpdateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	for i := range students {
		if fmt.Sprintf("%d", students[i].ID) == idStr {
			// 데이터가 들어온 경우에만 업데이트
			if input.Name != "" {
				students[i].Name = input.Name
			}
			if input.Age > 0 {
				students[i].Age = input.Age
			}
			if input.Dept != "" {
				students[i].Dept = input.Dept
			}

			c.JSON(http.StatusOK, gin.H{"status": "success", "data": students[i]})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "학생 없음"})
}

func DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	for i, s := range students {
		if fmt.Sprintf("%d", s.ID) == idStr {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "삭제 완료"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "학생 없음"})
}
