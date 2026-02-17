package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 기본 Gin 엔진 생성
	r := gin.Default()

	// 2. API 라우팅 그룹 설정 (버전 관리)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/students", GetStudents)    // 학생 목록 조회
		v1.POST("/students", CreateStudent) // 학생 등록
	}

	// 3. 서버 실행 (포트 8080)
	r.Run(":8080")
}
