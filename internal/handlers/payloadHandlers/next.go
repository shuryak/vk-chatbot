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

	skipCount := 0
	if p.Options != nil {
		skipCount = p.Options.SkipUsersCount
	}

	users, err := h.u.GetExceptOf(ctx, 2, skipCount, user.ID)
	if err != nil {
		return err
	}

	var msg *models.Message
	if len(users) != 0 {
		if len(users) == 1 {
			skipCount = 0
		} else {
			skipCount++
		}

		msg = models.NewTextMessage(reqMsg.PeerID, fmt.Sprintf("%s, %d лет, город %s.", users[0].Name, users[0].Age, users[0].City))
		msg.Keyboard = models.NewKeyboard(true).
			AddRow().
			AddButton("❤", models.PositiveColor, *models.NewPayload(models.LikeCommand, models.PayloadOptions{
				ShownUserID:    users[0].ID,
				SkipUsersCount: skipCount,
			})).
			AddButton("⛔", models.NegativeColor, *models.NewPayload(models.NextCommand, models.PayloadOptions{
				SkipUsersCount: skipCount,
			})).
			AddButtonWithCommandOnly("Снова он/она?", models.SecondaryColor, models.WhyISeeItCommand).
			AddButtonWithCommandOnly("Просто кнопка", models.SecondaryColor, models.NoCommand).
			AddRow().
			AddButtonWithCommandOnly("📕 Моя анкета", models.PrimaryColor, models.ShowCommand).
			AddButtonWithCommandOnly("👾 GitHub", models.SecondaryColor, models.GitHubCommand)
		msg.Attachment = &models.Attachment{PhotoID: users[0].PhotoID}
	} else {
		msg = models.NewTextMessage(reqMsg.PeerID, "🙈🙉 Нет доступных анкет.")
		msg.Keyboard = models.NewKeyboard(true).
			AddRow().
			AddButtonWithCommandOnly("📕 Моя анкета", models.PrimaryColor, models.ShowCommand).
			AddButtonWithCommandOnly("🔃 Попробовать снова", models.PositiveColor, models.NextCommand)
	}

	return h.messenger.Send(*msg)
}
