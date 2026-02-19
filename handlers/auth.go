package handlers

import (
	"go-study/database"
	"go-study/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// JWT를 서명할 때 사용할 비밀키 (실무에서는 환경변수로 관리해야 함)
var jwtKey = []byte("your_secret_key")

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. 유저 확인
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "아이디 또는 비밀번호가 틀렸습니다."})
		return
	}

	// 2. 비밀번호 검증
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "아이디 또는 비밀번호가 틀렸습니다."})
		return
	}

	// 3. 🆕 JWT 토큰 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24시간 후 만료
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "토큰 생성 실패"})
		return
	}

	// 4. 토큰 전달
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  tokenString,
	})
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
