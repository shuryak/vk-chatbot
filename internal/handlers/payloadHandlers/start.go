package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Start(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	msg := models.NewTextMessage(reqMsg.PeerID, "Окей, давай начнём :)")
	msg.Keyboard = models.NewKeyboard(true).
		AddRow().
		AddButtonWithCommandOnly("📌 Создать анкету", models.PositiveColor, models.SexCommand)

	err := h.messenger.Send(*msg)
	return err
}
