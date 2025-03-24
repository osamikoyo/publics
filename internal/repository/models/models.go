package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	MaxPeopleCount uint
	Date           time.Time
	Name           string
	Author         string
	Desc           string
	PeopleCount    int
}

type LoginRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	Username  string `gorm:"primaryKey"`
	Email     string
	Password  string
	Role      string
}

type Member struct {
	Event     Event
	User      User
	CreatedAt time.Time
	ID        uint `gorm:"primaryKey;autoIncrement"`
}
