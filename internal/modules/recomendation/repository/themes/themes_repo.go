package themes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/osamikoyo/publics/internal/modules/recomendation/entity"
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

func (repo *TopicRepository) CreateTopic(tc *entity.Topic, alikeID []uint) error {
	topic := tc.ToGraph()

	topic.DgraphType = "Topic"
	topicJSON, err := json.Marshal(topic)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		SetJson:   topicJSON,
		CommitNow: true,
	}

	ctx := context.Background()
	resp, err := repo.client.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return err
	}
	newTopicUID := resp.Uids["blank-0"]

	if len(alikeID) > 0 {
		query := fmt.Sprintf(`
			{
				findTopics(func: eq(id, %d)) @filter(type(Topic)) {
					uid
				}
			}
		`, alikeID[0])

		res, err := repo.client.NewTxn().Query(ctx, query)
		if err != nil {
			return err
		}

		var result struct {
			FindTopics []struct {
				UID string `json:"uid"`
			} `json:"findTopics"`
		}

		if err := json.Unmarshal(res.Json, &result); err != nil {
			return err
		}

		if len(result.FindTopics) > 0 {
			alikeTopicUID := result.FindTopics[0].UID

			edgeMutation := fmt.Sprintf(`
				<%s> <alike> <%s> .
			`, newTopicUID, alikeTopicUID)

			_, err = repo.client.NewTxn().Mutate(ctx, &api.Mutation{
				SetNquads: []byte(edgeMutation),
				CommitNow: true,
			})
			if err != nil {
				return err
			}
		}
	}
}
