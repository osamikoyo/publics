package repository

import (
	"fmt"

	"github.com/osamikoyo/publics/internal/repository/models"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *models.User) error
	GetBy(string, string) (*models.User, error)
	Login(req *models.LoginRequest) (*models.User, error)
}

type UserStorage struct {
	db     *gorm.DB
	logger *logger.Logger
}

func initUserRepository(db *gorm.DB, logger *logger.Logger) UserRepository {
	return &UserStorage{
		db:     db,
		logger: logger,
	}
}

func (storage *UserStorage) Register(user *models.User) error {
	res := storage.db.Create(user)

	if err := res.Error; err != nil {
		storage.logger.Error("cant do register db request", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return fmt.Errorf("cant do register db request: %v", err)
	}

	storage.logger.Info("successfully register user with", zapcore.Field{
		Key:    "username",
		String: user.Username,
	},
		zapcore.Field{
			Key:    "email",
			String: user.Email,
		})

	return nil
}

func (storage *UserStorage) GetBy(key, value string) (*models.User, error) {
	var User models.User

	res := storage.db.Model(&models.User{}).Where(fmt.Sprintf("%s = ?", key), value).Find(&User)
	if err := res.Error; err != nil {
		storage.logger.Error("cant do getby user request", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil, fmt.Errorf("cant do getby user request: %v", err)
	}

	return &User, nil
}

func (storage *UserStorage) Login(req *models.LoginRequest) (*models.User, error) {
	var (
		user models.User
		err  error
	)
	if req.Username == "" {
		err = storage.db.Model(&models.User{}).Where(&models.User{
			Email:    req.Email,
			Password: req.Password,
		}).Find(&user).Error
	} else if req.Email == "" {
		err = storage.db.Model(&models.User{}).Where(&models.User{
			Email:    req.Email,
			Password: req.Password,
		}).Find(&user).Error
	}

	if err != nil {
		storage.logger.Error("cant do login user request", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil, fmt.Errorf("cant login user: %v", err)
	}

	return &user, nil
}
