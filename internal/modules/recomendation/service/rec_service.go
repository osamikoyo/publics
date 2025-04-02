package service

import "github.com/osamikoyo/publics/internal/modules/event/entity"

type RecomendationService interface {
	GetRecs(uint) ([]entity.Event, error)
	AddFavTopic(uint, uint) error
	DeleteFavTopic(uint, uint) error
}
