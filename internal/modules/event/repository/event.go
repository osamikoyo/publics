package repository

import (
	"fmt"

	models "github.com/osamikoyo/publics/internal/modules/event/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type EventRepository interface {
	Add(*models.Event) error
	GetBy(string, string) ([]models.Event, error)
	Update(id uint, newEvent *models.Event) error
	Delete(id uint) error
}

type EventStorage struct {
	db     *gorm.DB
	logger *logger.Logger
}

func InitEventStorage(db *gorm.DB, logger *logger.Logger) (EventRepository, error) {
	return &EventStorage{
		db:     db,
		logger: logger,
	}, nil
}

func (e *EventStorage) Add(event *models.Event) error {
	return e.db.Create(event).Error
}

func (e *EventStorage) Update(id uint, newEvent *models.Event) error {
	res := e.db.Model(&models.Event{}).Where("id = ?", id).Find(newEvent)

	if err := res.Error; err != nil {
		e.logger.Error("cant do update req", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return fmt.Errorf("cant do update req: %v", err)
	}

	return nil
}

func (e *EventStorage) GetBy(key, value string) ([]models.Event, error) {
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

func (e *EventStorage) Delete(id uint) error {
	res := e.db.Where("id = ?", id).Delete(&models.Event{})
	if err := res.Error; err != nil {
		e.logger.Error("cant do delete request", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return fmt.Errorf("cant do delete request: %v", err)
	}

	return nil
}
