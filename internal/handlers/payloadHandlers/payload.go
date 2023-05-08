package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/logger"
)

type Handlers struct {
	messenger usecase.Messenger
	q         usecase.Questions
	um        usecase.ChatUsers
	u         usecase.Users
	s         usecase.Sympathy
	l         logger.Interface
}

func NewHandlers(messenger usecase.Messenger, q usecase.Questions, um usecase.ChatUsers, u usecase.Users, s usecase.Sympathy, l logger.Interface) *Handlers {
	return &Handlers{messenger, q, um, u, s, l}
}

func MessageFromContext(ctx context.Context) models.Message {
	return ctx.Value(models.MessageCtxKey).(models.Message)
}
