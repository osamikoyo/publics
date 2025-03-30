package checker

import (
	"time"

	"github.com/osamikoyo/publics/internal/modules/event/entity"
	member_entity "github.com/osamikoyo/publics/internal/modules/member/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Checker struct {
	db     *gorm.DB
	logger *logger.Logger
}

const TIME_FORMAT = "02.01.2006"

func Init(dsn string) (*Checker, error) {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &Checker{
		db:     db,
		logger: logger.Init(),
	}, nil
}

func (c *Checker) Check() error {
	c.logger.Info("starting check...")

	var events []entity.Event

	res := c.db.Where(&entity.Event{
		DateEnd: time.Now().Format(TIME_FORMAT),
	}).Find(&events)

	c.logger.Info("timeout events number", zapcore.Field{
		Key:     "number",
		Integer: int64(len(events)),
	})

	if err := res.Error; err != nil {
		c.logger.Error("error check", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})

		return err
	}

	res = c.db.Where(&entity.Event{
		DateEnd: time.Now().Format(TIME_FORMAT),
	}).Delete(&entity.Event{})

	c.logger.Info("deleted timeout events")

	for _, e := range events {
		res = c.db.Where(&member_entity.PublicMember{
			EventID: e.ID,
		}).Delete(&member_entity.PublicMember{})
	}

	c.logger.Info("deleted timeout members")

	return nil
}
