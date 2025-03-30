package entity

import (
	"github.com/osamikoyo/publics/internal/modules/user/entity"
)

type UserIndex struct {
	UserID uint
	User   entity.User `gorm:"foreighKey:UserID"`
	Topics []Topic     `gorm:"type:json"`
}
