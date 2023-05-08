package payloadHandlers

import (
	"context"
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
		_, err = h.s.Create(ctx, models.Sympathy{
			FirstUserID:  reqMsg.PeerID,
			SecondUserID: p.Options.ShownUserID,
			Reciprocity:  false,
		})
		if err != nil {
			h.l.Error("PayloadHandlers - Like - h.s.Create: %v", err)
			return err
		}
	} else {
		_, err = h.s.UpdateReciprocity(ctx, sym.ID, true)
		if err != nil {
			h.l.Error("PayloadHandlers - Like - h.s.UpdateReciprocity: %v", err)
			return err
		}
	}

	msgToSym := models.NewTextMessage(p.Options.ShownUserID, "Симпатия :)")

	h.messenger.Send(*msgToSym)

	h.Next(ctx, p)

	return nil
}
