package main

import (
	"go-study/database"
	"go-study/handlers"
	"go-study/middleware"
	"go-study/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.Student{})
	database.SeedData()

	r := gin.New()

	// 전역 미들웨어 적용
	r.Use(middleware.Logger())
	r.Use(gin.Recovery()) // 서버가 갑자기 죽는걸 방지해주는 기본 미들웨어

	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", handlers.GetStudents)
		v1.POST("/students", handlers.CreateStudent)
		v1.PATCH("/students/:id", handlers.UpdateStudent)
		v1.DELETE("/students/:id", middleware.Auth(), handlers.DeleteStudent)
	}

	r.Run(":8080")
}
