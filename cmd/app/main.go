package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/config"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"net/http"
)

func main() {
	var cfg *config.Config
	var err error

	configFilePath := flag.String("config", "", "Path to the config file")
	flag.Parse()
	if *configFilePath == "" {
		cfg, err = config.ParseEnv()
		if err != nil {
			log.Fatalf("config err: %v", err)
		}
		fmt.Println("Config file not provided. Only environment variables are used.")
	} else {
		cfg, err = config.ParseFileAndEnv(*configFilePath)
		if err != nil {
			log.Fatalf("config err: %v", err)
		}
		fmt.Printf("Using config file at: %s\n.", *configFilePath)
	}

	cb := callback.NewCallback()
	vk := vkapi.NewVKAPI(cfg.VK.Token)

	cb.ConfirmationKeys[cfg.VK.GroupID] = cfg.VK.ConfirmationKey

	cb.MessageNew(func(ctx context.Context, obj objects.MessageNewObject) {
		log.Println(obj.Message.Text)

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
	err = http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}
