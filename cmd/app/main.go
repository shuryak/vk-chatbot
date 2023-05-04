package main

import (
	"context"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/events"
	"log"
	"net/http"
)

func main() {
	cb := callback.NewCallback()

	cb.ConfirmationKeys[220319689] = "38b4c193"

	cb.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		log.Println(obj.Message.Text)
	})

	http.HandleFunc("/callback", cb.HandleFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
