package entity

import (
	"github.com/osamikoyo/publics/internal/modules/user/entity"
)

type UserIndex struct {
	UserID uint    `bson:"user_id"`
	Topics []Topic `bson:"topics"`
}
