package service

import (
	"github.com/osamikoyo/publics/internal/modules/member/entity"
	"github.com/osamikoyo/publics/internal/modules/member/repository"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type MemberService interface {
	Add(*entity.PublicMember) error
	Get(uint) ([]entity.PublicMember, error)
	CheckPerms(uint, uint) (bool, error)
	Delete(uint) error
}

type MemberServiceImpl struct {
	repo   repository.MemberRepository
	logger *logger.Logger
}

func (m *MemberServiceImpl) inject(repo repository.MemberRepository, logger *logger.Logger) *MemberServiceImpl {
	m.repo = repo
	m.logger = logger

	return m
}

func (m *MemberServiceImpl) Add(member *entity.PublicMember) error {
	if err := m.repo.Add(member); err != nil {
		m.logger.Error("error add member in repo", zapcore.Field{
			Key:     "user_id",
			Integer: int64(member.UserID),
		}, zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	return nil
}

func (m *MemberServiceImpl) Get(id uint) ([]entity.PublicMember, error) {
	members, err := m.repo.Get(id)
	if err != nil {
		m.logger.Info("error get members in repo", zapcore.Field{
			Key:     "req_id",
			Integer: int64(id),
		},
			zapcore.Field{
				Key:    "err",
				String: err.Error(),
			})
	}

	return nil, err
}

func (m *MemberServiceImpl) CheckPerms(EventID, UserID uint) (bool, error) {
	ok, err := m.repo.CheckPermitionDelete(EventID, UserID)
	if err != nil {
		m.logger.Error("error checkperms in repo", zapcore.Field{
			Key:     "event_id",
			Integer: int64(EventID),
		},
			zapcore.Field{
				Key:     "user_id",
				Integer: int64(UserID),
			},
			zapcore.Field{
				Key:    "err",
				String: err.Error(),
			})

		return false, err
	}

	return ok, nil
}

func (m *MemberServiceImpl) Delete(id uint) error {
	if err := m.repo.Delete(id); err != nil {
		m.logger.Error("error delete member in repo", zapcore.Field{
			Key:     "user_id",
			Integer: int64(id),
		}, zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	return nil
}
