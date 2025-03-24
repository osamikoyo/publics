package service

import (
	"fmt"

	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/internal/modules/user/repository"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Register(user *entity.User) error
	Login(req *entity.LoginRequest) (string, error)
	Auth(token string) bool
}

type userService struct {
	repo repository.UserRepository
	logger *logger.Logger
}

func Init(repo repository.UserRepository, logger *logger.Logger) UserService {
	return &userService{
		repo: repo,
		logger: logger,
	}
}

func (u *userService) Register(user *entity.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		u.logger.Error("cant register", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})

		return err
	}

	user.Password = string(password)

	return u.repo.Register(user)
}

func (u *userService) Login(req *entity.LoginRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		u.logger.Error("cant register", zapcore.Field{
			Key: "err",
			String: err.Error(),
		})

		return "", err
	}

	user, err := u.repo.Login(req)
	if err != nil || user == nil{
		return "", fmt.Errorf("cant auth: %v", err)
	}

	
}

func (u *userService) Auth(token string) bool {

}