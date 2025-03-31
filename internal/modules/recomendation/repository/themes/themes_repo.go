package themes

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/osamikoyo/publics/pkg/logger"
)

type TopicRepository struct {
  client *dgo.Dgraph
  logger *logger.Logger
}

func (repo *TopicRepository) Inject(client *dgo.Dgraph, logger *logger.Logger) *TopicRepository {
  repo.client = client
  repo.logger = logger

  return repo
}

func (repo *ThemesRepository) CreateTopic
