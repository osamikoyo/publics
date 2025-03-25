package service

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/osamikoyo/publics/internal/modules/user/entity"
	"github.com/osamikoyo/publics/internal/modules/user/interfaces/config"
	"github.com/osamikoyo/publics/internal/modules/user/repository"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *entity.User) error
	Login(req *entity.LoginRequest) (string, error)
	Auth(tkn string) (*Claims, error)
}

type UserPrivateService struct {
	repo   repository.UserRepository
	logger *logger.Logger
	key config.ServiceConfig
}

func (u *UserPrivateService) Inject(repo repository.UserRepository, logger *logger.Logger, cfg config.ServiceConfig) *UserPrivateService {
	u.repo = repo
	u.logger = logger
	u.key = cfg

	return u
}

func (u *UserPrivateService) Register(user *entity.User) error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("cant register", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	user.Password = string(password)

	return u.repo.Register(user)
}

func (u *UserPrivateService) Login(req *entity.LoginRequest) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("cant register", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return "", err
	}

	req.Password = string(password)

	user, err := u.repo.Login(req)
	if err != nil || user == nil {
		return "", fmt.Errorf("cant auth: %v", err)
	}

	return generateToken(user.ID, user.Username, u.key.GetKey())
}

func (u *UserPrivateService) Auth(tkn string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tkn, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(u.key.GetKey()), nil
	})

	if err != nil {
		u.logger.Error("cant auth user with", zapcore.Field{
			Key:    "token",
			String: tkn,
		})

		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token not valid")
}

