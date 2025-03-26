package servic

import (
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	"github.com/osamikoyo/publics/internal/modules/event/repository"
	"github.com/osamikoyo/publics/pkg/logger"
)

type EventService interface {
	Add(*entity.Event) error
	Update(uint, *entity.Event) error
	GetFirst(uint) (*entity.Event, error)
	Delte(uint) error
}

type EventServiceImpl struct {
	repo   repository.EventRepository
	logger *logger.Logger
}

func (e *EventServiceImpl) Inject(repo repository.EventRepository, logger *logger.Logger) *EventServiceImpl {
	e.repo = repo
	e.logger = logger

	return e
}

func (e *EventServiceImpl) Add(event *entity.Event) error {
	return e.repo.Add(event)
}

func (e *EventServiceImpl) Update(id uint, event *entity.Event) error {
	return e.repo.Update(id, event)
}
