package app

import (
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/config"
	"github.com/shuryak/vk-chatbot/internal/handlers"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"net/http"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	cb := callback.NewCallback(l)
	cb.ConfirmationKeys[cfg.VK.GroupID] = cfg.VK.ConfirmationKey
	vk := vkapi.NewVKAPI(l, cfg.VK.Token)

	h := handlers.NewHandlers(vk, l)

	cb.MessageNew(h.NewMessage)

	http.HandleFunc("/callback", cb.HandleFunc)

	fmt.Printf("Server running on %s.\n", cfg.Server.Port)
	err := http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}
