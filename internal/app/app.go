package app

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/config"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	cb := callback.NewCallback(l)
	cb.ConfirmationKeys[cfg.VK.GroupID] = cfg.VK.ConfirmationKey
	vk := vkapi.NewVKAPI(l, cfg.VK.Token)

	cb.MessageNew(func(ctx context.Context, obj objects.MessageNewObject) {
		l.Info("Message from %d received: %v", obj.Message.PeerID, obj.Message.Text)

		b := params.NewMessagesSendBuilder()
		b.Message("Hello!")
		b.RandomID(0)
		b.PeerID(obj.Message.PeerID)

		keyboard := objects.NewMessagesKeyboard(false)
		keyboard.AddRow()
		keyboard.AddTextButton("Привет :)", "hello-btn", objects.Primary)
		b.Keyboard(keyboard)

		//file, err := os.Open("test.jpg")
		//if err != nil {
		//	log.Fatal(err)
		//}
		//photo, err := vk.UploadMessagesPhoto(obj.Message.PeerID, file)
		//if err != nil {
		//	fmt.Println("Im here")
		//}
		//b.Attachment(photo[0].ToAttachment())

		resp, err := vk.MessagesSend(b.Params)
		log.Println(resp)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	fmt.Printf("Server running on %s.\n", cfg.Server.Port)
	err := http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}
