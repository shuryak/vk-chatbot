package handlers

import (
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
)

type Handlers struct {
	vkapi *vkapi.VKAPI
	l     logger.Interface
}

func NewHandlers(vkapi *vkapi.VKAPI, l logger.Interface) *Handlers {
	return &Handlers{vkapi, l}
}
