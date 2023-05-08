package payloadHandlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Like(ctx context.Context, p models.Payload) error {
	reqMsg := MessageFromContext(ctx)

	sym, err := h.s.GetByUserIDs(ctx, reqMsg.PeerID, p.Options.ShownUserID)
	if err != nil {
		h.l.Error("PayloadHandlers - Like - h.s.GetByUserIDs: %v", err)
		return err
	}

	if sym == nil {
		sym, err = h.s.Create(ctx, models.Sympathy{
			FirstUserID:  reqMsg.PeerID,
			SecondUserID: p.Options.ShownUserID,
			Reciprocity:  false,
		})
		if err != nil {
			h.l.Error("PayloadHandlers - Like - h.s.Create: %v", err)
			return err
		}
	} else {
		sym, err = h.s.UpdateReciprocity(ctx, sym.ID, true)
		if err != nil {
			h.l.Error("PayloadHandlers - Like - h.s.UpdateReciprocity: %v", err)
			return err
		}
	}

	user, err := h.u.GetByID(ctx, reqMsg.PeerID)
	if err != nil {
		return nil
	}

	msgToSym := models.NewTextMessage(
		p.Options.ShownUserID,
		fmt.Sprintf("😍 Ты понравился пользователю %s из города %s, %d лет", user.Name, user.City, user.Age),
	)
	msgToSym.Attachment = &models.Attachment{PhotoID: user.PhotoID}
	kbToSym := models.NewInlineKeyboard().
		AddRow().
		AddButton("💌 Ответить взаимностью", models.PrimaryColor, *models.NewPayload(
			models.ReciprocityCommand,
			models.PayloadOptions{
				ShownUserID: user.ID,
			},
		))
	msgToSym.Keyboard = kbToSym

	err = h.messenger.Send(*msgToSym)
	if err != nil {
		return err
	}

	err = h.Next(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
