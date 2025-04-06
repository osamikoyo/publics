package themes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/osamikoyo/publics/internal/modules/recomendation/entity"
	"github.com/osamikoyo/publics/pkg/logger"
	"go.uber.org/zap/zapcore"
)

type TopicRepository interface {
	CreateTopic(*entity.Topic, []uint)
	GetAlikeTopics(uint) []entity.Topic
}

type TopicStorage struct {
	client *dgo.Dgraph
	logger *logger.Logger
}

func (repo *TopicStorage) Inject(client *dgo.Dgraph, logger *logger.Logger) *TopicStorage {
	repo.client = client
	repo.logger = logger

	return repo
}

func (repo *TopicStorage) GetAlikeTopics(topicID uint) []entity.GraphTopic {
	ctx := context.Background()

	query := fmt.Sprintf(`
		query GetAlikeTopics($topicID: uint) {
			topic(func: uid($topicID)) {
				uid
				id
				text_explain
				desc
				dgraph.type
				
				~related_to {  
					uid
					id
					text_explain
					desc
					dgraph.type
				}
				
				related_to { 
					uid
					id
					text_explain
					desc
					dgraph.type
				}
			}
		}
	`)

	variables := map[string]string{
		"$topicID": fmt.Sprintf("%d", topicID),
	}

	resp, err := repo.client.NewTxn().QueryWithVars(ctx, query, variables)
	if err != nil {
		repo.logger.Error("Failed to execute Dgraph query: %v", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})
		return nil
	}

	var result struct {
		Topic []struct {
			UID          string              `json:"uid"`
			ID           uint                `json:"id"`
			TextExplain  string              `json:"text_explain"`
			Desc         string              `json:"desc"`
			DgraphType   string              `json:"dgraph.type"`
			RelatedTo    []entity.GraphTopic `json:"related_to"`
			RevRelatedTo []entity.GraphTopic `json:"~related_to"`
		} `json:"topic"`
	}

	if err := json.Unmarshal(resp.Json, &result); err != nil {
		repo.logger.Error("Failed to unmarshal Dgraph response", zapcore.Field{
			Key:    "err",
			String: err.Error(),
		})
		return nil
	}

	var alikeTopics []entity.GraphTopic
	if len(result.Topic) > 0 {
		alikeTopics = append(alikeTopics, result.Topic[0].RelatedTo...)
		alikeTopics = append(alikeTopics, result.Topic[0].RevRelatedTo...)
	}

	return alikeTopics
}

func (repo *TopicStorage) CreateTopic(tc *entity.Topic, alikeID []uint) error {
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

	return nil
}
