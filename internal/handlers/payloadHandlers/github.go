package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) GitHub(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)
	msg := models.NewTextMessage(reqMsg.PeerID, "Проект на GitHub: https://github.com/shuryak/vk-chatbot")
	msg.Keyboard = models.NewKeyboard(true).
		AddRow().
		AddButtonWithCommandOnly("📕 Моя анкета", models.PrimaryColor, models.ShowCommand)
	return h.messenger.Send(*msg)
}
