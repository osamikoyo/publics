package repository

import (
	"context"
	"time"

	"github.com/osamikoyo/publics/internal/modules/recomendation/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap/zapcore"
)

type RecomendationRepository struct {
	coll   *mongo.Collection
	logger *logger.Logger
	ctx    context.Context
}

func (repo *RecomendationRepository) Inject(coll *mongo.Collection, logger *logger.Logger) *RecomendationRepository {
	repo.coll = coll
	repo.logger = logger

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	repo.ctx = ctx

	return repo
}

func (repo *RecomendationRepository) GetUserFavouriteTopics(id uint) ([]entity.Topic, error) {
	var user entity.UserIndex

	filter := bson.M{
		"user_id": id,
	}

	res := repo.coll.FindOne(repo.ctx, filter)
	if err := res.Decode(&user); err != nil {
		repo.logger.Error("error get userIndex from mongo", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil, err
	}

	return user.Topics, nil
}
