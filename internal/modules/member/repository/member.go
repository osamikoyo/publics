package repository

import (
	"github.com/osamikoyo/publics/internal/modules/member/entity"
	"github.com/osamikoyo/publics/pkg/logger"
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

func (m *MemberRepository) Inject(db *gorm.DB, logger *logger.Logger) *MemberRepository {

}
