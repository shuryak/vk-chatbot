package app

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/shuryak/vk-chatbot/internal/config"
	"github.com/shuryak/vk-chatbot/internal/handlers"
	"github.com/shuryak/vk-chatbot/internal/handlers/payloadHandlers"
	"github.com/shuryak/vk-chatbot/internal/handlers/questionsHandlers"
	"github.com/shuryak/vk-chatbot/internal/models"
	"github.com/shuryak/vk-chatbot/internal/models/questions"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/internal/usecase/repo"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/postgres"
	"github.com/shuryak/vk-chatbot/pkg/vkapi"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/callback"
	"net/http"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Logger.Level)

	cb := callback.NewCallback(l)
	cb.ConfirmationKeys[cfg.VK.GroupID] = cfg.VK.ConfirmationKey
	vk := vkapi.NewVKAPI(l, cfg.VK.Token)

	pg, err := postgres.New(cfg.PG.URL, l, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	r := redis.NewClient(&redis.Options{
		Addr:       cfg.Redis.Host + cfg.Redis.Port,
		ClientName: cfg.Redis.Name,
		Password:   cfg.Redis.Password,
		DB:         0, // use default DB
	})

	uuc := usecase.NewUsersUseCase(repo.NewUsersRepo(pg))
	quc := usecase.NewQuestionsUseCase(repo.NewQuestionsRepo(r, 2*time.Minute))
	suc := usecase.NewSympathyUseCase(repo.NewSympathyRepo(pg))
	messenger := usecase.NewVKMessenger(vk)

	h := handlers.NewRegistry(quc, l)
	ph := payloadHandlers.NewHandlers(messenger, quc, usecase.NewVKUserManager(vk), *uuc, suc, l)
	qh := questionsHandlers.NewHandler(quc, *uuc, messenger)
	_ = h.RegisterPayloadHandler(models.StartCommand, ph.Start)
	_ = h.RegisterPayloadHandler(models.SexCommand, ph.Sex)
	_ = h.RegisterPayloadHandler(models.CreateCommand, ph.Create)
	_ = h.RegisterPayloadHandler(models.ShowCommand, ph.Show)
	_ = h.RegisterPayloadHandler(models.NextCommand, ph.Next)
	_ = h.RegisterPayloadHandler(models.LikeCommand, ph.Like)
	_ = h.RegisterPayloadHandler(models.DislikeCommand, ph.Dislike)
	_ = h.RegisterPayloadHandlerForMany(ph.Change, models.CityCommand, models.NameCommand, models.AgeCommand)
	_ = h.RegisterQuestionHandlerForMany(qh.Edit, questions.CityQuestion, questions.NameQuestion, questions.AgeQuestion)

	cb.MessageNew(h.Handle)

	http.HandleFunc("/callback", cb.HandleFunc)

	fmt.Printf("Server running on %s.\n", cfg.Server.Port)
	err = http.ListenAndServe(cfg.Server.Port, nil)
	if err != nil {
		panic(err)
	}
}
