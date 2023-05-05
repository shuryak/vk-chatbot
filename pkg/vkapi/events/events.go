package events

import (
	"context"
	"encoding/json"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/transport"
)

type EventType string

// https://vk.com/dev/groups_events
const (
	EventConfirmation = "confirmation"
	EventMessageNew   = "message_new"
	EventMessageReply = "message_reply"
	EventMessageEdit  = "message_edit"
)

type GroupEvent struct {
	Type    EventType       `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	V       string          `json:"v"`
	Secret  string          `json:"secret"`
}

func NewFuncList(l logger.Interface) *FuncList {
	return &FuncList{l: l}
}

type FuncList struct {
	messageNew []func(context.Context, objects.MessageNewObject)
	eventsList []EventType

	l logger.Interface

	goroutine bool
}

func (fl *FuncList) Handler(ctx context.Context, e GroupEvent) error {
	ctx = context.WithValue(ctx, transport.GroupIDKey, e.GroupID)
	ctx = context.WithValue(ctx, transport.EventIDKey, e.EventID)
	ctx = context.WithValue(ctx, transport.EventVersionKey, e.V)
	switch e.Type {
	case EventMessageNew:
		var obj objects.MessageNewObject
		if err := json.Unmarshal(e.Object, &obj); err != nil {
			fl.l.Error("FuncList - EventMessageNew - json.Unmarshal: %v", err)
			return err
		}

		for _, f := range fl.messageNew {
			f := f // A local copy of the function to use it inside the anonymous function.

			if fl.goroutine {
				go func() { f(ctx, obj) }()
			} else {
				f(ctx, obj)
			}
		}
	}

	return nil
}

func (fl *FuncList) MessageNew(f func(ctx context.Context, object objects.MessageNewObject)) {
	fl.messageNew = append(fl.messageNew, f)
	fl.eventsList = append(fl.eventsList, EventMessageNew)
}
