package payload

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
)

func (h *Handlers) Change(ctx context.Context, p models.Payload) error {
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
	}

	err := h.q.Set(ctx, reqMsg.PeerID, qType)
	if err != nil {
		return err
	}

	msg := models.NewTextMessage(reqMsg.PeerID, msgText)

	err = h.messenger.Send(*msg)
	return err
}
