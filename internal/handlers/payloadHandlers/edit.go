package payloadHandlers

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
)

func (h *Handlers) Edit(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	var msgText string
	var qType questions.QuestionType

	switch *p.Command {
	case models.CityCommand:
		msgText = "✏️ Введи свой город:"
		qType = questions.CityQuestion
	case models.AgeCommand:
		msgText = "✏️ Введи свой возраст:"
		qType = questions.AgeQuestion
	case models.NameCommand:
		msgText = "✏️ Введи своё имя:"
		qType = questions.NameQuestion
	case models.InterestedInCommand:
		msgText = "✅ Будем искать "
		if p.Options.InterestedIn == "girls" {
			msgText += "девушек"
		} else if p.Options.InterestedIn == "boys" {
			msgText += "парней"
		}
		user := models.User{
			ID:           reqMsg.PeerID,
			InterestedIn: p.Options.InterestedIn,
		}
		_, err := h.u.Update(ctx, user)
		if err != nil {
			return err
		}
		qType = questions.NoQuestion
	}

	msg := models.NewTextMessage(reqMsg.PeerID, msgText)

	if qType != questions.NoQuestion {
		err := h.q.Set(ctx, reqMsg.PeerID, qType)
		if err != nil {
			return err
		}
	} else {
		msg.Keyboard = models.NewKeyboard(true).
			AddRow().
			AddButtonWithCommandOnly("📕 Моя анкета", models.PrimaryColor, models.ShowCommand)
	}

	return h.messenger.Send(*msg)
}
