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
	database.DB.AutoMigrate(&models.Student{}, &models.User{})
	database.SeedData()

	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.Static("/uploads", "./uploads")

	v1 := r.Group("/api/v1")
	{
		// 회원가입 API 추가
		v1.POST("/register", handlers.Register)
		v1.POST("/login", handlers.Login)

		//프로필 사진 업로드 API
		v1.POST("/students/:id/upload", middleware.Auth(), handlers.UploadProfile)

		v1.GET("/students", handlers.GetStudents)
		v1.POST("/students", handlers.CreateStudent)
		v1.PATCH("/students/:id", middleware.Auth(), handlers.UpdateStudent)
		v1.DELETE("/students/:id", middleware.Auth(), handlers.DeleteStudent)
	}
	r.MaxMultipartMemory = 8 << 20
	r.Run(":8080")
}
