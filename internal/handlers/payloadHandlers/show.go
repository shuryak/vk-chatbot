package payloadHandlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Show(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	user, err := h.u.GetByID(ctx, reqMsg.PeerID)
	if err != nil {
		return err
	}

	interestedInText := "девушки"
	negInterestedIn := "boys"
	negInterestedInText := "парни"
	if user.InterestedIn == "boys" {
		interestedInText, negInterestedInText = negInterestedInText, interestedInText
		negInterestedIn = "girls"
	}

	msg := models.NewTextMessage(
		reqMsg.PeerID,
		fmt.Sprintf("Имя: %s. Город: %s. Возраст: %d. Интересуют: %s.\n\nПроверь правильность данных. Если всё 👌, нажми зелёную кнопку и переходи к просмотру анкет.", user.Name, user.City, user.Age, interestedInText),
	)
	msg.Keyboard = models.NewInlineKeyboard().
		AddRow().
		AddButton("✅ Всё верно", models.PositiveColor, *models.NewPayloadWithCommandOnly(models.NextCommand)).
		AddRow().
		AddButton("👑 Изменить имя", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.NameCommand)).
		AddRow().
		AddButton("🏙️ Изменить город", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.CityCommand)).
		AddRow().
		AddButton("5️⃣ Изменить возраст", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.AgeCommand)).
		AddRow().
		AddButton("⚧ Интересуют "+negInterestedInText, models.SecondaryColor, *models.NewPayload(models.InterestedInCommand, models.PayloadOptions{
			InterestedIn: negInterestedIn,
		}))

	msg.Attachment = &models.Attachment{PhotoID: user.PhotoID}

	err = h.messenger.Send(*msg)
	return err
}
