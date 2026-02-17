package main

import (
	"go-study/database"
	"go-study/handlers"
	"go-study/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//DB 초기화
	database.InitDB()
	//테이블 자동 생성 (Migration)
	database.DB.AutoMigrate(&models.Student{})
	//더미 데이터 채우기
	database.SeedData()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", handlers.GetStudents)
		v1.POST("/students", handlers.CreateStudent)
		v1.PATCH("/students/:id", handlers.UpdateStudent)
		v1.DELETE("/students/:id", handlers.DeleteStudent)
	}
	r.Run(":8080")
}
