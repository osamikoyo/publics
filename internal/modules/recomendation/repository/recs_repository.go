package repository

import (
	"github.com/osamikoyo/publics/internal/modules/recomendation/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/gorm"
)

type RecomendationRepository struct {
	db     *gorm.DB
	logger *logger.Logger
}

func (repo *RecomendationRepository) Inject(db *gorm.DB, logger *logger.Logger) *RecomendationRepository {
	repo.db = db
	repo.logger = logger

	return repo
}

func (repo *RecomendationRepository) GetUserFavouriteTopics(id uint) ([]entity.Topic, error) {

}
