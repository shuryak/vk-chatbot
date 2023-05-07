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

type PayloadHandlerFunc func(ctx context.Context, p models.Payload) error

type QuestionHandlerFunc func(ctx context.Context) error

type Registry struct {
	payloadHandlers  map[models.ButtonCommand]PayloadHandlerFunc
	questionHandlers map[questions.QuestionType]QuestionHandlerFunc
	q                usecase.Questions
	l                logger.Interface
}

func NewRegistry(q usecase.Questions, l logger.Interface) *Registry {
	return &Registry{
		make(map[models.ButtonCommand]PayloadHandlerFunc),
		make(map[questions.QuestionType]QuestionHandlerFunc),
		q,
		l,
	}
}

func (h *Registry) RegisterPayloadHandler(cmd models.ButtonCommand, handler PayloadHandlerFunc) error {
	if _, ok := h.payloadHandlers[cmd]; ok {
		return fmt.Errorf("%s payloadHandlers handler already registered", cmd)
	}
	h.payloadHandlers[cmd] = handler
	return nil
}

func (h *Registry) RegisterPayloadHandlerForMany(handler PayloadHandlerFunc, cmds ...models.ButtonCommand) error {
	for _, cmd := range cmds {
		err := h.RegisterPayloadHandler(cmd, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Registry) UnregisterPayloadHandler(cmd models.ButtonCommand) error {
	if _, ok := h.payloadHandlers[cmd]; ok {
		delete(h.payloadHandlers, cmd)
		return nil
	}
	return fmt.Errorf("%s payloadHandlers handler is not registered", cmd)
}

func (h *Registry) RegisterQuestionHandler(q questions.QuestionType, handler QuestionHandlerFunc) error {
	if _, ok := h.questionHandlers[q]; ok {
		return fmt.Errorf("%s question handler already registered", q)
	}
	h.questionHandlers[q] = handler
	return nil
}

func (h *Registry) RegisterQuestionHandlerForMany(handler QuestionHandlerFunc, questions ...questions.QuestionType) error {
	for _, q := range questions {
		err := h.RegisterQuestionHandler(q, handler)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Registry) UnregisterQuestionHandler(q questions.QuestionType) error {
	if _, ok := h.questionHandlers[q]; ok {
		delete(h.questionHandlers, q)
		return nil
	}
	return fmt.Errorf("%s question handler is not registered", q)
}

func (h *Registry) Handle(ctx context.Context, obj objects.MessageNewObject) {
	h.l.Info("Message from %d received: %v. Payload: %v", obj.Message.PeerID, obj.Message.Text, obj.Message.Payload)

	payload, err := models.UnmarshalPayload(obj.Message.Payload)
	if err != nil {
		h.l.Error("Registry - Handle - models.UnmarshalPayload: %v", err)
	}

	msg := models.NewTextMessage(obj.Message.PeerID, obj.Message.Text)

	if payload.Command != nil {
		if _, ok := h.payloadHandlers[*payload.Command]; ok {
			err := h.payloadHandlers[*payload.Command](ContextWithMessage(ctx, *msg), payload)
			if err != nil {
				h.l.Error("Registry - Handle - h.payloadHandlers: %v", err)
			}
			return
		}
	} else {
		q, err := h.q.Get(ctx, obj.Message.PeerID)
		if err != nil {
			h.l.Error("Registry - Handle - h.qm.Get: %v", err)
			return
		}
		err = h.questionHandlers[q](ContextWithQuestion(ContextWithMessage(ctx, *msg), q))
		if err != nil {
			h.l.Error("Registry - Handle - h.questionHandlers: %v", err)
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
