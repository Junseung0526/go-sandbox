package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student 구조체: JSON 태그와 검증(binding) 태그 포함
type Student struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gt=0"`
	Dept string `json:"dept"`
}

// 가상의 데이터베이스 (메모리 슬라이스)
// 가상의 데이터베이스 (메모리 슬라이스)
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

// [GET] 전체 학생 목록 조회
func GetStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   students,
	})
}

// [POST] 학생 추가 (에러 보강 버전)
func CreateStudent(c *gin.Context) {
	var newStudent Student

	// 1. JSON 바인딩 및 필수값 검증
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "입력 형식이 잘못되었거나 필수 값이 누락되었습니다.",
			"error":   err.Error(),
		})
		return
	}

	// 2. ID 중복 체크 (비즈니스 로직)
	for _, s := range students {
		if s.ID == newStudent.ID {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "error",
				"message": "이미 존재하는 학생 ID입니다.",
			})
			return
		}
	}

	// 3. 데이터 저장
	students = append(students, newStudent)
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newStudent,
	})
}
