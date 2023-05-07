package payload

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/logger"
)

type Handlers struct {
	messenger usecase.Messenger
	uuc       usecase.UsersUseCase
	qr        usecase.QuestionsRepo
	l         logger.Interface
}

func NewHandlers(messenger usecase.Messenger, uuc usecase.UsersUseCase, qr usecase.QuestionsRepo, l logger.Interface) *Handlers {
	return &Handlers{
		messenger: messenger,
		uuc:       uuc,
		qr:        qr,
		l:         l,
	}
}

func MessageFromContext(ctx context.Context) models.Message {
	return ctx.Value(models.MessageCtxKey).(models.Message)
}
