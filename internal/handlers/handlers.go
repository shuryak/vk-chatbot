package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
)

type PayloadHandler func(ctx context.Context, p models.Payload) error

type QuestionHandler func(ctx context.Context) error

type PayloadHandlers struct {
	payloadHandlers  map[models.ButtonCommand]PayloadHandler
	questionHandlers map[questions.QuestionType]QuestionHandler
	q                usecase.Questions
	l                logger.Interface
}

func NewPayloadHandlers(q usecase.Questions, l logger.Interface) *PayloadHandlers {
	return &PayloadHandlers{
		make(map[models.ButtonCommand]PayloadHandler),
		make(map[questions.QuestionType]QuestionHandler),
		q,
		l,
	}
}

func (h *PayloadHandlers) RegisterPayloadHandler(cmd models.ButtonCommand, handler PayloadHandler) error {
	if _, ok := h.payloadHandlers[cmd]; ok {
		return fmt.Errorf("%s payload handler already registered", cmd)
	}
	h.payloadHandlers[cmd] = handler
	return nil
}

func (h *PayloadHandlers) RegisterPayloadHandlerForMany(handler PayloadHandler, cmds ...models.ButtonCommand) error {
	for _, cmd := range cmds {
		err := h.RegisterPayloadHandler(cmd, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *PayloadHandlers) UnregisterPayloadHandler(cmd models.ButtonCommand) error {
	if _, ok := h.payloadHandlers[cmd]; ok {
		delete(h.payloadHandlers, cmd)
		return nil
	}
	return fmt.Errorf("%s payload handler is not registered", cmd)
}

func (h *PayloadHandlers) RegisterQuestionHandler(q questions.QuestionType, handler QuestionHandler) error {
	if _, ok := h.questionHandlers[q]; ok {
		return fmt.Errorf("%s question handler already registered", q)
	}
	h.questionHandlers[q] = handler
	return nil
}

func (h *PayloadHandlers) RegisterQuestionHandlerForMany(handler QuestionHandler, questions ...questions.QuestionType) error {
	for _, q := range questions {
		err := h.RegisterQuestionHandler(q, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *PayloadHandlers) UnregisterQuestionHandler(q questions.QuestionType) error {
	if _, ok := h.questionHandlers[q]; ok {
		delete(h.questionHandlers, q)
		return nil
	}
	return fmt.Errorf("%s question handler is not registered", q)
}

func (h *PayloadHandlers) Handle(ctx context.Context, obj objects.MessageNewObject) {
	h.l.Info("Message from %d received: %v. Payload: %v", obj.Message.PeerID, obj.Message.Text, obj.Message.Payload)

	payload, err := models.UnmarshalPayload(obj.Message.Payload)
	if err != nil {
		h.l.Error("PayloadHandlers - Handle - models.UnmarshalPayload: %v", err)
	}

	msg := models.NewTextMessage(obj.Message.PeerID, obj.Message.Text)

	if payload.Command != nil {
		if _, ok := h.payloadHandlers[*payload.Command]; ok {
			err := h.payloadHandlers[*payload.Command](ContextWithMessage(ctx, *msg), payload)
			if err != nil {
				h.l.Error("PayloadHandlers - Handle - h.payloadHandlers: %v", err)
			}
			return
		}
	} else {
		q, err := h.q.Get(ctx, obj.Message.PeerID)
		if err != nil {
			h.l.Error("PayloadHandlers - Handle - h.qm.Get: %v", err)
			return
		}
		err = h.questionHandlers[q](ContextWithQuestion(ContextWithMessage(ctx, *msg), q))
		if err != nil {
			h.l.Error("PayloadHandlers - Handle - h.questionHandlers: %v", err)
		}
		return
	}
}

func ContextWithMessage(parent context.Context, msg models.Message) context.Context {
	return context.WithValue(parent, models.MessageCtxKey, msg)
}

func ContextWithQuestion(parent context.Context, q questions.QuestionType) context.Context {
	return context.WithValue(parent, models.QuestionCtxKey, q)
}
