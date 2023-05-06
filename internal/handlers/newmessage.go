package handlers

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/payload"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/params"
	"log"
	"strconv"
	"time"
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
	case payload.StartCommand:
		b.Message("Окей, давай начнём :)")
		keyboard := objects.NewMessagesKeyboard(true)
		keyboard.AddRow()
		keyboard.AddTextButton("📌 Создать анкету", payload.New(payload.SexCommand), objects.Positive)
		b.Keyboard(keyboard)
	case payload.SexCommand:
		b.Message("Кого будем искать?")
		keyboard := objects.NewMessagesKeyboard(true)
		keyboard.AddRow()

		keyboard.AddTextButton("👩 Девушки", payload.Payload{
			Command: payload.AboutCommand,
			Options: payload.Options{
				InterestedIn: "girls",
			},
		}, objects.Negative)

		keyboard.AddTextButton("👨 Парни", payload.Payload{
			Command: payload.AboutCommand,
			Options: payload.Options{
				InterestedIn: "boys",
			},
		}, objects.Primary)

		b.Keyboard(keyboard)
	case payload.AboutCommand:
		usersGetBuilder := params.NewUsersGetBuilder()
		usersGetBuilder.UserIDs([]string{strconv.Itoa(obj.Message.PeerID)})
		usersGetBuilder.Fields([]string{"photo_id, city, bdate"})

		users, err := h.vkapi.UsersGet(usersGetBuilder.Params)
		if err != nil {
			h.l.Error("Handlers - NewMessage - h.vkapi.UsersGet: %v", err)
		}

		t, err := time.Parse("2.1.2006", users[0].Bdate)
		if err != nil {
			fmt.Println(err)
			return
		}
		now := time.Now()
		age := now.Year() - t.Year()
		if now.YearDay() < t.YearDay() {
			age--
		}

		keyboard := objects.NewMessagesKeyboardInline()
		keyboard.AddRow()
		keyboard.AddTextButton("✅ Всё верно", payload.Payload{
			Command: payload.SaveCommand,
			Options: payload.Options{
				InterestedIn: pl.Options.InterestedIn,
				Name:         users[0].FirstName,
				Age:          age,
				City:         users[0].City.Title,
			},
		}, objects.Positive)
		keyboard.AddRow()
		keyboard.AddTextButton("🏙️ Изменить город", payload.New(payload.CityCommand), objects.Primary)
		keyboard.AddRow()
		keyboard.AddTextButton("5️⃣ Изменить возраст", payload.New(payload.AgeCommand), objects.Primary)

		b.Keyboard(keyboard)

		b.Message(fmt.Sprintf("%s, город %s. Возраст: %d. Интересуют: %s", users[0].FirstName, users[0].City.Title, age, pl.Options.InterestedIn))
		b.Attachment("photo" + users[0].PhotoId)
	case payload.CityCommand:
		b.Message("🏙️ Введи свой город:")
		err := h.qr.Set(ctx, obj.Message.PeerID, entities.CityQuestion)
		if err != nil {
			log.Fatal(err) // TODO: handle error
		}
	case payload.SaveCommand:
		_, err := h.uuc.Create(ctx, entities.User{
			VKID:         obj.Message.PeerID,
			PhotoURL:     "",
			Name:         pl.Options.Name,
			City:         pl.Options.City,
			InterestedIn: pl.Options.InterestedIn,
		})
		if err != nil {
			return
		}
		b.Message("Анкета сохранена!")
	default:
		q, err := h.qr.Get(ctx, obj.Message.PeerID)
		if err != nil {
			log.Fatal(err)
		}

		if q == entities.CityQuestion {
			b.Message(fmt.Sprintf("✅ Город установлен на %s", obj.Message.Text))
		}
	}

	resp, err := h.vkapi.MessagesSend(b.Params)
	log.Println(resp)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}
