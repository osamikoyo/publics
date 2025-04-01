package themes

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
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

func (repo *ThemesRepository) CreateTopic(tc *entity.Topic, alikeID []uint) error {
	topic := 

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
	resp, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return err
	}

	// Получаем UID новой темы
	newTopicUID := resp.Uids["blank-0"] // или другой ключ, если задан

	// 2. Если есть alikeIDs, добавляем связи
	if len(alikeIDs) > 0 {
		// Формируем запрос для поиска UID тем по их ID
		query := fmt.Sprintf(`
			{
				findTopics(func: eq(id, %d)) @filter(type(Topic)) {
					uid
				}
			}
		`, alikeID[0]) // Пример для одного ID (можно расширить)

		res, err := dgraphClient.NewTxn().Query(ctx, query)
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

			_, err = dgraphClient.NewTxn().Mutate(ctx, &api.Mutation{
				SetNquads: []byte(edgeMutation),
				CommitNow: true,
			})
			if err != nil {
				return err
			}
		}
	}
}
