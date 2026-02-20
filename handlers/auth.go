package handlers

import (
	"go-study/database"
	"go-study/models"
	"net/http"
	"os" // 🆕 추가
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	// 🆕 환경 변수에서 시크릿 키 로드
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		jwtKey = []byte("your_secret_key")
	}

	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "인증 실패"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 생성 실패"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": tokenString})
}

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
