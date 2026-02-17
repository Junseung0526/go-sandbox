package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt 자동 생성
	Name       string `json:"name" binding:"required"`
	Age        int    `json:"age" binding:"required,gt=0"`
	Dept       string `json:"dept"`
}

type UpdateStudentInput struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"omitempty,gt=0"`
	Dept string `json:"dept"`
}
