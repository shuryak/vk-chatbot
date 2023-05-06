package repo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"strconv"
	"time"
)

type QuestionsRepo struct {
	*redis.Client
	lifetime time.Duration
}

func NewQuestionsRepo(r *redis.Client, lifetime time.Duration) *QuestionsRepo {
	return &QuestionsRepo{r, lifetime}
}

// Check for implementation
var _ usecase.QuestionsRepo = (*QuestionsRepo)(nil)

func (qr QuestionsRepo) Set(ctx context.Context, VKID int, q entities.QuestionType) error {
	return qr.Client.Set(ctx, strconv.Itoa(VKID), q, qr.lifetime).Err()
}

func (qr QuestionsRepo) Get(ctx context.Context, VKID int) (entities.QuestionType, error) {
	v, err := qr.Client.Get(ctx, strconv.Itoa(VKID)).Result()
	if err == redis.Nil {
		return entities.NoQuestion, nil
	}
	if err != nil {
		return entities.NoQuestion, err
	}
	return entities.QuestionType(v), nil
}
