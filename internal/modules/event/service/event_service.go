package servic

import (
	"time"

	"github.com/osamikoyo/publics/internal/modules/event/entity"
	"github.com/osamikoyo/publics/internal/modules/event/repository"
	"github.com/osamikoyo/publics/pkg/logger"
)

type EventService interface {
	Add(*entity.Event) error
	Update(uint, *entity.Event) error
	GetBy(string, string) ([]entity.Event, error)
	Delete(uint) error
}

type EventServiceImpl struct {
	repo   repository.EventRepository
	logger *logger.Logger
}

func (e *EventServiceImpl) Inject(repo repository.EventRepository, logger *logger.Logger) EventService {
	e.repo = repo
	e.logger = logger

	return e
}

const TIME_FORMAT = "02.01.2006"

func (e *EventServiceImpl) Add(event *entity.Event) error {
	event.DateEnd = time.Now().Format(TIME_FORMAT)

	return e.repo.Add(event)
}

func (e *EventServiceImpl) Update(id uint, event *entity.Event) error {
	return e.repo.Update(id, event)
}

func (e *EventServiceImpl) GetBy(key string, value string) ([]entity.Event, error) {
	return e.repo.GetBy(key, value)
}

func (e *EventServiceImpl) Delete(id uint) error {
	return e.repo.Delete(id)
}
