package payloadHandlers

import (
	"context"
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
	_, err = h.u.Create(ctx, models.User{
		ID:           reqUser.ID,
		PhotoID:      reqUser.PhotoID,
		Name:         reqUser.Name,
		Age:          age,
		City:         reqUser.City,
		InterestedIn: p.Options.InterestedIn,
		Activated:    &f,
	})
	if err != nil {
		return err
	}

	return h.Show(ctx, p)
}
