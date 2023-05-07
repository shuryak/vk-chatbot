package payloadHandlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
	"time"
)

func (h *Handlers) Create(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	reqUser, err := h.um.GetByID(reqMsg.PeerID)
	if err != nil {
		return err
	}

	var age int
	if reqUser.BirthDate != nil {
		now := time.Now()
		age = now.Year() - reqUser.BirthDate.Year()
		if now.YearDay() < reqUser.BirthDate.YearDay() {
			age--
		}
	}

	f := false
	user, err := h.u.Create(ctx, models.User{
		ID:           reqUser.ID,
		PhotoID:      reqUser.PhotoID,
		Name:         reqUser.Name,
		Age:          age,
		City:         reqUser.City,
		InterestedIn: p.Options.InterestedIn,
		Activated:    &f,
	})

	var interestedIn string
	if user.InterestedIn == "girls" {
		interestedIn = "девушки"
	} else {
		interestedIn = "парни"
	}

	msg := models.NewTextMessage(
		reqMsg.PeerID,
		fmt.Sprintf("Имя: %s. Город: %s. Возраст: %d. Интересуют: %s.\n\nПроверь правильность данных. Если всё 👌, нажми зелёную кнопку и переходи к просмотру анкет.", user.Name, user.City, user.Age, interestedIn),
	)
	msg.Keyboard = models.NewInlineKeyboard().
		AddRow().
		AddButton("✅ Всё верно", models.PositiveColor, *models.NewPayloadWithCommandOnly(models.SaveCommand)).
		AddRow().
		AddButton("👑 Изменить имя", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.NameCommand)).
		AddRow().
		AddButton("🏙️ Изменить город", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.CityCommand)).
		AddRow().
		AddButton("5️⃣ Изменить возраст", models.SecondaryColor, *models.NewPayloadWithCommandOnly(models.AgeCommand))
	msg.Attachment = &models.Attachment{PhotoID: user.PhotoID}

	err = h.messenger.Send(*msg)
	return err
}
