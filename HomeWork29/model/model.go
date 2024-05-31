package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string
	Lastname   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}