package database

import (
	"fmt"
	"go-study/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("students.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("상세 에러 내용: %v\n", err)
		panic("데이터베이스 연결 실패!")
	}

	fmt.Println("데이터베이스 연결 성공!")
}

func SeedData() {
	var count int64
	DB.Model(&models.Student{}).Count(&count)

	// 데이터가 없을 때만 생성 (중복 방지)
	if count == 0 {
		dummyStudents := []models.Student{
			{Name: "Kim Junseung", Age: 24, Dept: "Smart Software"},
			{Name: "Min-ji", Age: 22, Dept: "Computer Science"},
			{Name: "Seung-woo", Age: 25, Dept: "AI Engineering"},
			{Name: "Ji-hyeon", Age: 23, Dept: "Smart Software"},
			{Name: "Do-yun", Age: 21, Dept: "Electronic Engineering"},
			{Name: "Ha-rin", Age: 22, Dept: "Nursing"},
		}

		for _, s := range dummyStudents {
			DB.Create(&s)
		}
		fmt.Println("더미 데이터 시딩 완료! 🌱")
	}
}
