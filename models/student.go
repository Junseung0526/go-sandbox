package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required,gt=0"`
	Dept         string `json:"dept"`
	ProfileImage string `json:"profile_image"`
}

type UpdateStudentInput struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"omitempty,gt=0"`
	Dept string `json:"dept"`
}
