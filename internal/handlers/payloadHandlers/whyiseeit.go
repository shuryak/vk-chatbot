package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) WhyISeeIt(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)
	msg := models.NewTextMessage(
		reqMsg.PeerID,
		"Подбор анкет ❗цикличен❗ и основан на том, в поиске пользователей какого пола ты заинтересован.",
	)
	msg.Keyboard = models.NewKeyboard(true).
		AddRow().
		AddButtonWithCommandOnly("📕 Моя анкета", models.PrimaryColor, models.ShowCommand)
	return h.messenger.Send(*msg)
}
