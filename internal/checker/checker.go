package checker

import (
	"github.com/osamikoyo/publics/internal/modules/event/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Checker struct{
  db *gorm.DB
  logger *logger.Logger
}

func Init(dsn string) (*Checker, error) {
  db, err := gorm.Open(sqlite.Open(dsn))
  if err != nil{
    return nil, err
  }

  return &Checker{
    db: db,
    logger: logger.Init(),
  }, nil
} 

func (c *Checker) Check() error {
  var events []entity.Event

  res := c.db.Where(&entity.Event{
    Date
  })
}
