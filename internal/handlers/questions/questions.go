package questions

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
	"github.com/shuryak/vk-chatbot/internal/usecase"
)

type Handler struct {
	q         usecase.Questions
	messenger usecase.Messenger
}

func NewHandler(q usecase.Questions, messenger usecase.Messenger) *Handler {
	return &Handler{q, messenger}
}

func QuestionFromContext(ctx context.Context) questions.QuestionType {
	return ctx.Value(models.QuestionCtxKey).(questions.QuestionType)
}
