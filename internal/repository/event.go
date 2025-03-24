package repository

import (
	"fmt"

	"github.com/osamikoyo/publics/internal/config"
	"github.com/osamikoyo/publics/internal/repository/models"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type EventRepo struct {
	db     *gorm.DB
	logger *logger.Logger
}

func initEventRepo(cfg *config.Config, logger *logger.Logger) (*EventRepo, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBpath))
	if err != nil {
		logger.Error("cant connect to db", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil, fmt.Errorf("cant open db: %v", err)
	}

	return &EventRepo{
		db:     db,
		logger: logger,
	}, nil
}

func (e *EventRepo) Add(event *models.Event) error {
	return e.db.Create(event).Error
}

func (e *EventRepo) GetBy(key, value string) ([]models.Event, error) {
	var events []models.Event

	res := e.db.Where(fmt.Sprintf("%s = %s", key, value)).Find(&events)
	if res.Error != nil {
		e.logger.Error("cant do getby db request", zapcore.Field{
			Key:    "err",
			String: res.Error.Error(),
		})

		return nil, fmt.Errorf("cant do getby req: %v", res.Error)
	}

	return events, nil
}

func (e *EventRepo) Delete(id uint) error {
	return e.db.Where("ID = ?", id).Delete(&models.Event{}).Error
}
