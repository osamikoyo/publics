package service

import (
	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *entity.User) error
	Login(req *entity.LoginRequest) (string, error)
	Auth(token string) bool
}

type userService struct {
	db     *gorm.DB
	logger *logger.Logger
}

func Init(db *gorm.DB, logger *logger.Logger) {

}
