package service

import (
	"github.com/osamikoyo/publics/internal/repository/models"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *models.User) error
	Login(req *models.LoginRequest) (string, error)
	Auth(token string) bool
}

type userService struct {
	db     *gorm.DB
	logger *logger.Logger
}

func Init(db *gorm.DB, logger *logger.Logger) {

}
