package handlers

import (
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
)

type Handlers struct {
	vkapi *vkapi.VKAPI
	uuc   usecase.UsersUseCase
	qr    usecase.QuestionsRepo
	l     logger.Interface
}

func NewHandlers(vkapi *vkapi.VKAPI, uuc usecase.UsersUseCase, qr usecase.QuestionsRepo, l logger.Interface) *Handlers {
	return &Handlers{vkapi, uuc, qr, l}
}
