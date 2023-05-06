package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"strconv"
)

func (h *Handlers) NewMessage(ctx context.Context, obj objects.MessageNewObject) {
	h.l.Info("Message from %d received: %v", obj.Message.PeerID, obj.Message.Text)

	b2 := params.NewUsersGetBuilder()
	b2.UserIDs([]string{strconv.Itoa(obj.Message.PeerID)})
	b2.Fields([]string{"photo_id"})

	resp, err := h.vkapi.UsersGet(b2.Params)

	b := params.NewMessagesSendBuilder()
	b.Message(fmt.Sprintf("Привет, %s с аватаркой: %s", resp[0].FirstName, resp[0].PhotoId))
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	resp2, err := h.vkapi.MessagesSend(b.Params)
	log.Println(resp2)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp2)
}
