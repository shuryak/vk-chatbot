package usecase

import (
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"time"
)

type VKMessenger struct {
	*vkapi.VKAPI
}

func NewVKMessenger(vkapi *vkapi.VKAPI) *VKMessenger {
	return &VKMessenger{vkapi}
}

// Check for implementation
var _ Messenger = (*VKMessenger)(nil)

func (vk VKMessenger) Send(msg models.Message) error {
	builder := params.NewMessagesSendBuilder().
		PeerID(msg.PeerID).
		Message(msg.Text).
		RandomID(time.Now().Nanosecond())

	if msg.Keyboard != nil {
		kb := VKKeyboard(*msg.Keyboard)
		builder.Keyboard(kb)
	}
	if msg.Attachment != nil {
		builder.Attachment("photo" + msg.Attachment.PhotoID)
	}

	_, err := vk.MessagesSend(builder.Params)
	return err
}

func VKKeyboard(keyboard models.Keyboard) (VKKeyboard *objects.MessagesKeyboard) {
	if keyboard.Inline {
		VKKeyboard = objects.NewMessagesKeyboardInline()
	} else {
		VKKeyboard = objects.NewMessagesKeyboard(keyboard.OneTime)
	}

	for i := 0; i < len(keyboard.Buttons); i++ {
		VKKeyboard.AddRow()
		for j := 0; j < len(keyboard.Buttons[i]); j++ {
			btn := keyboard.Buttons[i][j]
			VKKeyboard.AddTextButton(btn.Text, btn.Command, string(btn.Color))
		}
	}

	return
}
