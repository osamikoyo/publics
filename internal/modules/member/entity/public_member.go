package entity

import (
	"time"
)

type PublicMember struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	EventID   uint
	UserID    uint
	CreatedAt time.Time
	RoleOn    string
	Username  string
}
