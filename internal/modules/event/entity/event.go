package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Regular        bool
	DateEnd        string
	CreatedAt      time.Time
	Author         string
	MaxMemberCount uint
	MemberCount    uint
	Title          string
	Desc           string
	Topic          string
}
