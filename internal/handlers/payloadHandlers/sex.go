package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Sex(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	msg := models.NewTextMessage(reqMsg.PeerID, "Кого будем искать?")
	msg.Keyboard = models.NewKeyboard(true).
		AddRow().
		AddButton("👩 Девушки", models.NegativeColor, *models.NewPayload(models.CreateCommand, models.PayloadOptions{
			InterestedIn: "girls",
		})).
		AddRow().
		AddButton("👨 Парни", models.PrimaryColor, *models.NewPayload(models.CreateCommand, models.PayloadOptions{
			InterestedIn: "boys",
		}))

	return h.messenger.Send(*msg)
}
