package questions

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/handlers/payload"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
	"strconv"
)

func (h *Handler) Edit(ctx context.Context) error {
	reqMsg := payload.MessageFromContext(ctx)
	q := QuestionFromContext(ctx)

	user := models.User{
		ID: reqMsg.PeerID,
	}
	var msgText string

	switch q {
	case questions.CityQuestion:
		user.City = reqMsg.Text
		msgText = "✅ Город установлен на " + user.City
	case questions.AgeQuestion:
		age, err := strconv.Atoi(reqMsg.Text)
		if err != nil {
			return err
		}
		user.Age = age
		msgText = "✅ Возраст установлен на " + strconv.Itoa(user.Age)
	case questions.NameQuestion:
		user.Name = reqMsg.Text
		msgText = "✅ Имя установлено на " + user.Name
	}

	_, err := h.u.Update(ctx, user)
	if err != nil {
		return err
	}

	msg := models.NewTextMessage(reqMsg.PeerID, msgText)
	msg.Keyboard = models.NewKeyboard(true).
		AddRow().
		AddButtonWithCommandOnly("📕 Посмотреть анкету", models.PrimaryColor, models.ShowCommand)

	err = h.messenger.Send(*msg)
	return err
}
