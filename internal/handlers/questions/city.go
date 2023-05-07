package questions

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/handlers/payload"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handler) Edit(ctx context.Context) error {
	reqMsg := payload.MessageFromContext(ctx)
	q := QuestionFromContext(ctx)

	msg := models.NewTextMessage(reqMsg.PeerID, string(q))

	err := h.messenger.Send(*msg)
	return err
}
