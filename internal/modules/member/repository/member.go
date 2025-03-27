package repository

import (
	"github.com/osamikoyo/publics/internal/modules/member/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Add(*entity.PublicMember) error
	Get(uint) ([]entity.PublicMember, error)
	Delete(*entity.PublicMember) error
}

type MemberStorage struct {
	db     *gorm.DB
	logger *logger.Logger
}

func (repo *MemberStorage) Inject(db *gorm.DB, logger *logger.Logger) *MemberRepository {
	repo.db = db
	repo.logger = logger

	return repo
}

func (repo *MemberStorage) Add(member *entity.PublicMember) error {
	res := repo.db.Create(member)
	if err := res.Error; err != nil {
		repo.logger.Error("cant add member to db", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	return nil
}

func (repo *MemberStorage) Get(id uint) ([]entity.PublicMember, error) {
	var members []entity.PublicMember

	res := repo.db.Model(&entity.PublicMember{}).Where(&entity.PublicMember{
		EventID: id,
	}).Find(&members)

	if err := res.Error; err != nil {
		repo.logger.Error("cant get members from db", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil, err
	}

	return members, nil
}

func (repo *MemberStorage) Delete(id uint) error {
	res := repo.db.Delete(&entity.PublicMember{}, id)

	if err := res.Error; err != nil {
		repo.logger.Error("cant delete member from db", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return nil
	}

	return nil
}
