package questionsHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
	"github.com/shuryak/vk-chatbot/internal/usecase"
)

type Handler struct {
	q         usecase.Questions
	u         usecase.Users
	messenger usecase.Messenger
}

func NewHandler(q usecase.Questions, u usecase.Users, messenger usecase.Messenger) *Handler {
	return &Handler{q, u, messenger}
}

func QuestionFromContext(ctx context.Context) questions.QuestionType {
	return ctx.Value(models.QuestionCtxKey).(questions.QuestionType)
}
