package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/payload"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"strconv"
)

func (h *Handlers) NewMessage(ctx context.Context, obj objects.MessageNewObject) {
	h.l.Info("Message from %d received: %v. Payload: %v", obj.Message.PeerID, obj.Message.Text, obj.Message.Payload)

	pl, err := payload.Unmarshal(obj.Message.Payload)
	if err != nil {
		h.l.Error("Handlers - NewMessage - payload.Unmarshal: %v", err)
	}

	b := params.NewMessagesSendBuilder()
	b.RandomID(0)
	b.PeerID(obj.Message.PeerID)

	switch pl.Command {
	case payload.Start:
		keyboard := objects.NewMessagesKeyboard(true)
		keyboard.AddRow()
		keyboard.AddTextButton("📌 Создать анкету", payload.New(payload.Create), objects.Positive)
		b.Keyboard(keyboard)
		b.Message("Окей, давай начнём :)")
	case payload.Create:
		usersGetBuilder := params.NewUsersGetBuilder()
		usersGetBuilder.UserIDs([]string{strconv.Itoa(obj.Message.PeerID)})
		usersGetBuilder.Fields([]string{"photo_id, city"})

		users, err := h.vkapi.UsersGet(usersGetBuilder.Params)
		if err != nil {
			h.l.Error("Handlers - NewMessage - h.vkapi.UsersGet: %v", err)
		}

		b.Message(fmt.Sprintf("%s, город %v.", users[0].FirstName, users[0].City.Title))
		b.Attachment("photo" + users[0].PhotoId)
	}

	resp, err := h.vkapi.MessagesSend(b.Params)
	log.Println(resp)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
