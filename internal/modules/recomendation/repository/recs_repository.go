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

func (repo *RecomendationRepository) AddFavoutiteTopic(id uint, topic *entity.Topic) error {
	filter := bson.M{
		"user_id": id,
	}

	update := bson.M{
		"$push": bson.M{
			"topics": topic,
		},
	}

	if err := repo.coll.FindOneAndUpdate(repo.ctx, filter, update).Err(); err != nil {
		repo.logger.Error("error add favourite topic", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	return nil
}

func (repo *RecomendationRepository) DeleteFavouriteTopic(id uint, topicID uint) error {
	filter := bson.M{
		"user_id": id,
	}

	update := bson.M{
		"$pull": bson.M{
			"id": topicID,
		},
	}

	res := repo.coll.FindOneAndUpdate(repo.ctx, filter, update)
	if err := res.Err(); err != nil {
		repo.logger.Error("error delete favourite topic", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	return nil
}
