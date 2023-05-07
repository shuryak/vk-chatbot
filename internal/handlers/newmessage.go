package handlers

//func (h *PayloadHandlers) NewMessage(ctx context.Context, obj objects.MessageNewObject) {
//	//h.l.Info("Message from %d received: %v. Payload: %v", obj.Message.PeerID, obj.Message.Text, obj.Message.Payload)
//
//	//pl, err := payload.Unmarshal(obj.Message.Payload)
//	//if err != nil {
//	//h.l.Error("PayloadHandlers - NewMessage - payloads.Unmarshal: %v", err)
//	//}
//
//	h.Handle(ctx)
//
//	b := params.NewMessagesSendBuilder()
//	b.RandomID(0)
//	b.PeerID(obj.Message.PeerID)
//
//	h.RegisterHandler(models.StartCommand)
//
//	switch pl.Command {
//	case payload.StartCommand:
//		b.Message("Окей, давай начнём :)")
//		keyboard := objects.NewMessagesKeyboard(true)
//		keyboard.AddRow()
//		keyboard.AddTextButton("📌 Создать анкету", payload.New(payload.SexCommand), objects.Positive)
//		b.Keyboard(keyboard)
//	case payload.SexCommand:
//		b.Message("Кого будем искать?")
//		keyboard := objects.NewMessagesKeyboard(true)
//		keyboard.AddRow()
//
//		keyboard.AddTextButton("👩 Девушки", payload.Payload{
//			Command: payload.AboutCommand,
//			Options: payload.Options{
//				InterestedIn: "girls",
//			},
//		}, objects.Negative)
//
//		keyboard.AddTextButton("👨 Парни", payload.Payload{
//			Command: payload.AboutCommand,
//			Options: payload.Options{
//				InterestedIn: "boys",
//			},
//		}, objects.Primary)
//
//		b.Keyboard(keyboard)
//	case payload.AboutCommand:
//		usersGetBuilder := params.NewUsersGetBuilder()
//		usersGetBuilder.UserIDs([]string{strconv.Itoa(obj.Message.PeerID)})
//		usersGetBuilder.Fields([]string{"photo_id, city, bdate"})
//
//		users, err := h.vkapi.UsersGet(usersGetBuilder.Params)
//		if err != nil {
//			h.l.Error("PayloadHandlers - NewMessage - h.vkapi.UsersGet: %v", err)
//		}
//
//		t, err := time.Parse("2.1.2006", users[0].Bdate)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		now := time.Now()
//		age := now.Year() - t.Year()
//		if now.YearDay() < t.YearDay() {
//			age--
//		}
//
//		f := false
//		user, err := h.uuc.Create(ctx, entities.User{
//			VKID:         obj.Message.PeerID,
//			PhotoURL:     users[0].PhotoId,
//			Name:         users[0].FirstName,
//			Age:          age,
//			City:         users[0].City.Title,
//			InterestedIn: pl.Options.InterestedIn,
//			Activated:    &f,
//		})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		keyboard := objects.NewMessagesKeyboardInline()
//		keyboard.AddRow()
//		keyboard.AddTextButton("✅ Всё верно", payload.New(payload.SaveCommand), objects.Positive)
//		keyboard.AddRow()
//		keyboard.AddTextButton("👑 Изменить имя", payload.New(payload.NameCommand), objects.Secondary)
//		keyboard.AddRow()
//		keyboard.AddTextButton("🏙️ Изменить город", payload.New(payload.CityCommand), objects.Secondary)
//		keyboard.AddRow()
//		keyboard.AddTextButton("5️⃣ Изменить возраст", payload.New(payload.AgeCommand), objects.Secondary)
//
//		b.Keyboard(keyboard)
//
//		b.Message(fmt.Sprintf("%s, город %s. Возраст: %d. Интересуют: %s", user.Name, user.City, user.Age, user.InterestedIn))
//		b.Attachment("photo" + users[0].PhotoId)
//	case payload.CityCommand:
//		b.Message("✏️ Введи свой город:")
//		err := h.qr.Set(ctx, obj.Message.PeerID, entities.CityQuestion)
//		if err != nil {
//			log.Fatal(err) // TODO: handle error
//		}
//	case payload.AgeCommand:
//		b.Message("✏️ Введи свой возраст:")
//		err := h.qr.Set(ctx, obj.Message.PeerID, entities.AgeQuestion)
//		if err != nil {
//			log.Fatal(err) // TODO: handle error
//		}
//	case payload.NameCommand:
//		b.Message("✏️ Введи своё имя:")
//		err := h.qr.Set(ctx, obj.Message.PeerID, entities.NameQuestion)
//		if err != nil {
//			log.Fatal(err) // TODO: handle error
//		}
//	case payload.SaveCommand:
//		user, err := h.uuc.GetByVKID(ctx, obj.Message.PeerID)
//		if err != nil {
//			log.Fatal(err) // TODO: handler error
//		}
//
//		keyboard := objects.NewMessagesKeyboardInline()
//		keyboard.AddRow()
//		keyboard.AddTextButton("✅ Всё верно", payload.Payload{
//			Command: payload.SaveCommand,
//			Options: payload.Options{
//				InterestedIn: user.InterestedIn,
//				Name:         user.Name,
//				Age:          user.Age,
//				City:         user.City,
//			},
//		}, objects.Positive)
//		keyboard.AddRow()
//		keyboard.AddTextButton("👑 Изменить имя", payload.New(payload.NameCommand), objects.Secondary)
//		keyboard.AddRow()
//		keyboard.AddTextButton("🏙️ Изменить город", payload.New(payload.CityCommand), objects.Secondary)
//		keyboard.AddRow()
//		keyboard.AddTextButton("5️⃣ Изменить возраст", payload.New(payload.AgeCommand), objects.Secondary)
//
//		b.Keyboard(keyboard)
//
//		b.Message(fmt.Sprintf("%s, город %s. Возраст: %d. Интересуют: %s", user.Name, user.City, user.Age, user.InterestedIn))
//		b.Attachment("photo" + user.PhotoURL)
//	default:
//		q, err := h.qr.Get(ctx, obj.Message.PeerID)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		keyboard := objects.NewMessagesKeyboard(true)
//		keyboard.AddRow()
//		keyboard.AddTextButton("💾 Сохранить изменения", payload.New(payload.SaveCommand), objects.Positive)
//
//		switch q {
//		case entities.CityQuestion:
//			user, err := h.uuc.Update(ctx, entities.User{
//				VKID: obj.Message.PeerID,
//				City: obj.Message.Text,
//			})
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			b.Message(fmt.Sprintf("✅ Город установлен на %s", user.City))
//			b.Keyboard(keyboard)
//		case entities.AgeQuestion:
//			age, err := strconv.Atoi(obj.Message.Text)
//			if err != nil {
//				b.Message("Неверный формат")
//			}
//
//			user, err := h.uuc.Update(ctx, entities.User{
//				VKID: obj.Message.PeerID,
//				Age:  age,
//			})
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			b.Message(fmt.Sprintf("✅ Возраст установлен на %d", user.Age))
//			b.Keyboard(keyboard)
//		case entities.NameQuestion:
//			user, err := h.uuc.Update(ctx, entities.User{
//				VKID: obj.Message.PeerID,
//				Name: obj.Message.Text,
//			})
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			b.Message(fmt.Sprintf("✅ Имя установлено на %s", user.Name))
//			b.Keyboard(keyboard)
//		}
//	}
//
//	resp, err := h.vkapi.MessagesSend(b.Params)
//	log.Println(resp)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(resp)
//}
