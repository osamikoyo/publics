package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID             uint64 `gorm:"primaryKey;autoIncrement"`
	Regular        bool
	Date           time.Time
	Author         string
	MaxMemberCount uint
	MemberCount    uint
	Title          string
	Desc           string
}
