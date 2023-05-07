package usecase

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/models"
)

type (
	Users interface {
		Create(ctx context.Context, u entities.User) (*entities.User, error)
		GetByVKID(ctx context.Context, VKID int) (*entities.User, error)
		Update(ctx context.Context, u entities.User) (*entities.User, error)
	}

	Messenger interface {
		Send(msg models.Message) error
	}

	UsersRepo interface {
		Create(ctx context.Context, u entities.User) (*entities.User, error)
		GetByVKID(ctx context.Context, VKID int) (*entities.User, error)
		Update(ctx context.Context, VKID int, columns Columns) (*entities.User, error)
		// TODO: Full CRUD
	}

	SympathyRepo interface {
		Create(ctx context.Context, s entities.Sympathy) (*entities.Sympathy, error)
		GetByUserIDs(ctx context.Context, firstUserID, secondUserID int) (*entities.Sympathy, error)
		UpdateReciprocity(ctx context.Context, id int, reciprocity bool) (*entities.Sympathy, error)
		// TODO: Full CRUD
	}

	QuestionsRepo interface {
		Set(ctx context.Context, VKID int, q entities.QuestionType) error
		Get(ctx context.Context, VKID int) (entities.QuestionType, error)
	}
)
