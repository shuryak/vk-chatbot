package main

import (
	"context"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"net/http"
)

func main() {
	cb := callback.NewCallback()
	vk := vkapi.NewVKAPI("vk1.a.GK7hvGI58YFlR45psBWriGWNKULnmRZQJTew8ViIUecJe4F8lUHr7PxRZhPMpHDBn1uMyDz5wbtfzt-gdTFZhZMTECI2TyzupZVtBMEfu3-hyeLBvJY-h0xoVJopwK3gThA7X1lPilxNurjhUTqjWMaMrMvLx28VOkspHHGNYrGQak7IFI0RJNk4kbEzl202Vb_lbcD09uG1A_lg3IlF0Q")

	cb.ConfirmationKeys[220319689] = "8f504305"

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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
