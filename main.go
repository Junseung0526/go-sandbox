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
	// User 모델도 자동 마이그레이션에 추가
	database.DB.AutoMigrate(&models.Student{}, &models.User{})
	database.SeedData()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		// 회원가입 API 추가
		v1.POST("/register", handlers.Register)
		v1.POST("/login", handlers.Login)

		v1.GET("/students", handlers.GetStudents)
		v1.POST("/students", handlers.CreateStudent)
		v1.PATCH("/students/:id", middleware.Auth(), handlers.UpdateStudent)
		v1.DELETE("/students/:id", middleware.Auth(), handlers.DeleteStudent)
	}

	r.Run(":8080")
}
