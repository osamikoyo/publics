package service

import (
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	topicrepo "github.com/osamikoyo/publics/internal/modules/recomendation/repository/themes"
)

type RecomendationService interface {
	GetRecs(uint) ([]entity.Event, error)
	AddFavTopic(uint, uint) error
	DeleteFavTopic(uint, uint) error
}

type RecomendationServiceImpl struct {
	topicRepo topicrepo.TopicRepository
}
