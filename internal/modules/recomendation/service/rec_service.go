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
	topicConnectingsStorage map[*selfentity.Element]float32
}

func (r *RecomendationServiceImpl) Inject(topicRepo topicrepo.TopicRepository, recsRepo repository.RecomendationRepository,
	eventRepo eventrepo.EventRepository) *RecomendationServiceImpl {
	r.recsRepo = recsRepo
	r.topicRepo = topicRepo
	r.eventRepo = eventRepo

	r.topicConnectingsStorage = make(map[*selfentity.Element]float32)

	return r
}

func (r *RecomendationServiceImpl) getProcentOf(full int, element *selfentity.Element) {
	self := r.topicConnectingsStorage[element]
	if self != 0 {
		r.topicConnectingsStorage[element] = float32(full / 100)
	} else {
		r.topicConnectingsStorage[element] = float32(full / (int(r.topicConnectingsStorage[element] + 1)))
	}
}

func (r *RecomendationServiceImpl) GetRecs() ([]entity.Event, error) {

}
