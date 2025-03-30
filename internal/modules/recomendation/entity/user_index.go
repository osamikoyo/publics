package entity

import (
	"github.com/osamikoyo/publics/internal/modules/user/entity"
)

type UserIndex struct {
	User   entity.User
	Topics []Topic `gorm:"type:json"`
}
