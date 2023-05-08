package payloadHandlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Next(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	t := true
	user, err := h.u.Update(ctx, models.User{
		ID:        reqMsg.PeerID,
		Activated: &t,
	})
	if err != nil {
		return err
	}

	msg := models.NewTextMessage(reqMsg.PeerID, "Теперь в БД: ")
	if *user.Activated {
		msg.Text += "activated = true"
	} else {
		msg.Text += "activated = false"
	}

	users, err := h.u.GetExceptOf(ctx, 10, 0, reqMsg.PeerID)
	if err != nil {
		return err
	}

	if len(users) != 0 {
		msg := models.NewTextMessage(reqMsg.PeerID, fmt.Sprintf("%s, %d лет, город %s.", users[0].Name, users[0].Age, users[0].City))
		msg.Keyboard = models.NewKeyboard(true).
			AddRow().
			AddButton("❤", models.PositiveColor, *models.NewPayload(models.LikeCommand, models.PayloadOptions{
				ShownUserID: users[0].ID,
			})).
			AddButton("⛔", models.NegativeColor, *models.NewPayload(models.DislikeCommand, models.PayloadOptions{
				ShownUserID: users[0].ID,
			}))
		msg.Attachment = &models.Attachment{PhotoID: users[0].PhotoID}
		h.messenger.Send(*msg)
	}

	return err
}
