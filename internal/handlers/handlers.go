package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
)

type PayloadHandler func(ctx context.Context, p models.Payload) error

type PayloadHandlers struct {
	handlers map[models.ButtonCommand]PayloadHandler
	l        logger.Interface
}

func NewHandlers(l logger.Interface) *PayloadHandlers {
	return &PayloadHandlers{make(map[models.ButtonCommand]PayloadHandler), l}
}

func (h *PayloadHandlers) RegisterHandler(cmd models.ButtonCommand, ph PayloadHandler) error {
	if _, ok := h.handlers[cmd]; ok {
		return fmt.Errorf("%s handler already registered", cmd)
	}
	h.handlers[cmd] = ph
	return nil
}

func (h *PayloadHandlers) UnregisterHandler(cmd models.ButtonCommand) error {
	if _, ok := h.handlers[cmd]; ok {
		delete(h.handlers, cmd)
		return nil
	}
	return fmt.Errorf("%s handler is not registered", cmd)
}

func (h *PayloadHandlers) Handle(ctx context.Context, obj objects.MessageNewObject) {
	h.l.Info("Message from %d received: %v. Payload: %v", obj.Message.PeerID, obj.Message.Text, obj.Message.Payload)

	payload, err := models.UnmarshalPayload(obj.Message.Payload)
	if err != nil {
		h.l.Error("PayloadHandlers - Handle - models.UnmarshalPayload: %v", err)
	}

	msg := models.NewTextMessage(obj.Message.PeerID, obj.Message.Text)

	if _, ok := h.handlers[payload.Command]; ok {
		err := h.handlers[payload.Command](ContextWithMessage(ctx, *msg), payload)
		if err != nil {
			h.l.Error("PayloadHandlers - Handle - models.UnmarshalPayload: %v", err)
		}
	}
}

func ContextWithMessage(parent context.Context, msg models.Message) context.Context {
	return context.WithValue(parent, models.MessageCtxKey, msg)
}
