package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Dislike(ctx context.Context, p models.Payload) error {
	//reqMsg := MessageFromContext(ctx)

	return h.Next(ctx, p)
}
