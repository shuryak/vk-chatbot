package payload

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/models"
)

func (h *Handlers) Start(ctx context.Context, p models.Payload) error {
	msg := MessageFromContext(ctx)
	fmt.Printf("Сообщение: %s\n; Payload: %s", msg.Text, p.Command)

	err := h.messenger.Send(*models.NewTextMessage(msg.PeerID, "Hello World!"))
	if err != nil {
		return err
	}

	return nil
}
