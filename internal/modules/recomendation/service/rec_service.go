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

const LAST_PARENTS = 3

type RecomendationServiceImpl struct {
	topicRepo               topicrepo.TopicRepository
	recsRepo                repository.RecomendationRepository
	eventRepo               eventrepo.EventRepository
	arrayOfUID              []selfentity.UID
	topicSelfStorage        map[selfentity.UID]*selfentity.Element
	topicConnectingsStorage map[selfentity.UID]float32
}

func (r *RecomendationServiceImpl) Inject(topicRepo topicrepo.TopicRepository, recsRepo repository.RecomendationRepository,
	eventRepo eventrepo.EventRepository) *RecomendationServiceImpl {
	r.recsRepo = recsRepo
	r.topicRepo = topicRepo
	r.eventRepo = eventRepo

	r.topicSelfStorage = make(map[selfentity.UID]*selfentity.Element)
	r.topicConnectingsStorage = make(map[selfentity.UID]float32)

	return r
}

func (r *RecomendationServiceImpl) incrementProcentOf(full int, element *selfentity.Element) {
	r.arrayOfUID = append(r.arrayOfUID, element.ID)
	r.topicSelfStorage[element.ID] = element

	if r.topicConnectingsStorage[element.ID] != 0 {
		r.topicConnectingsStorage[element.ID] = float32(((r.topicConnectingsStorage[element.ID] + 1) * 100) / float32(full))
	} else {
		r.topicConnectingsStorage[element.ID] = float32((1 * 100) / full)
	}
}

func (r *RecomendationServiceImpl) getProcentOf(topic string) float32 {
	for _, el := range r.arrayOfUID {
		if r.topicSelfStorage[el].Self.Topic == topic {
			return r.topicConnectingsStorage[el]
		}
	}

	return 0
}

func (r *RecomendationServiceImpl) AddFavTopic(userID, topicID uint) error {

}

func (r *RecomendationServiceImpl) GetRecs() ([]entity.Event, error) {

}
