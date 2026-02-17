package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", GetStudents)
		v1.POST("/students", CreateStudent)
		v1.PATCH("/students/:id", UpdateStudent)
		v1.DELETE("/students/:id", DeleteStudent)
	}

	r.Run(":8080")
}
