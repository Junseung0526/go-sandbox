package handlers

import (
	"go-study/database"
	"go-study/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 비밀번호 암호화 (Cost: 14)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	// DB 저장
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "이미 존재하는 사용자입니다."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "회원가입 성공!"})
}

// [POST] 로그인
func Login(c *gin.Context) {
	var input models.User
	var user models.User

	// 1. 입력 데이터 바인딩
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "error": err.Error()})
		return
	}

	// 2. DB에서 유저 찾기
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "아이디 또는 비밀번호가 틀렸습니다."})
		return
	}

	// 3. 비밀번호 비교 (bcrypt)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "아이디 또는 비밀번호가 틀렸습니다."})
		return
	}

	// 4. 성공 응답 (나중에는 여기서 JWT 토큰을 발급합니다)
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "로그인 성공!",
		"user": gin.H{
			"username": user.Username,
		},
	})
}
