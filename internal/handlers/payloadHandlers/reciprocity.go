package payloadHandlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Reciprocity(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	msg := models.NewTextMessage(
		reqMsg.PeerID,
		fmt.Sprintf("💑 Отличная пара! Взаимная симпатия с @id%d", p.Options.ShownUserID),
	)
	err := h.messenger.Send(*msg)
	if err != nil {
		return err
	}

	msgToSym := models.NewTextMessage(
		p.Options.ShownUserID,
		fmt.Sprintf("💑 Отличная пара! Взаимная симпатия с @id%d", reqMsg.PeerID),
	)
	err = h.messenger.Send(*msgToSym)
	if err != nil {
		return err
	}

	return nil
}
