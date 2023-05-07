package usecase

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
)

type QuestionsUseCase struct {
	repo QuestionsRepo
}

func NewQuestionsUseCase(qr QuestionsRepo) *QuestionsUseCase {
	return &QuestionsUseCase{qr}
}

// Check for implementation
var _ Questions = (*QuestionsUseCase)(nil)

func (qm *QuestionsUseCase) Set(ctx context.Context, ID int, q questions.QuestionType) error {
	err := qm.repo.Set(ctx, ID, entities.QuestionType(q))
	return err
}

func (qm *QuestionsUseCase) Get(ctx context.Context, ID int) (questions.QuestionType, error) {
	q, err := qm.repo.Get(ctx, ID)
	return questions.QuestionType(q), err
}

func (qm *QuestionsUseCase) Delete(ctx context.Context, ID int) error {
	err := qm.repo.Delete(ctx, ID)
	return err
}
