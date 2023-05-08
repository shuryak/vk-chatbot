package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Dislike(ctx context.Context, p models.Payload) error {
	//reqMsg := MessageFromContext(ctx)

	h.Next(ctx, p)
	return nil
}
