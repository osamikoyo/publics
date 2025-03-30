package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Email    string
	Username string
	Password string
	Role     string
}

