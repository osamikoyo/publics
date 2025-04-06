package service

import (
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	eventrepo "github.com/osamikoyo/publics/internal/modules/event/repository"
	selfentity "github.com/osamikoyo/publics/internal/modules/recomendation/entity"
	"github.com/osamikoyo/publics/internal/modules/recomendation/repository"
	topicrepo "github.com/osamikoyo/publics/internal/modules/recomendation/repository/themes"
)

type RecomendationService interface {
	GetRecs(uint) ([]entity.Event, error)
	AddFavTopic(uint, uint) error
	DeleteFavTopic(uint, uint) error
}

type RecomendationServiceImpl struct {
	topicRepo               topicrepo.TopicRepository
	recsRepo                repository.RecomendationRepository
	eventRepo               eventrepo.EventRepository
	topicSelfStorage        map[selfentity.UID]*selfentity.Element
	topicConnectingsStorage map[selfentity.UID]float32
}

func (r *RecomendationServiceImpl) Inject(topicRepo topicrepo.TopicRepository, recsRepo repository.RecomendationRepository,
	eventRepo eventrepo.EventRepository) *RecomendationServiceImpl {
	r.recsRepo = recsRepo
	r.topicRepo = topicRepo
	r.eventRepo = eventRepo

	r.topicConnectingsStorage = make(map[selfentity.UID]float32)
	r.topicSelfStorage = make(map[selfentity.UID]*selfentity.Element)

	return r
}

func (r *RecomendationServiceImpl) getProcentOf(full int, element *selfentity.Element) {

}

func (r *RecomendationServiceImpl) GetRecs() ([]entity.Event, error) {

}
