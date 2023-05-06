package usecase

import (
	"context"
	"github.com/shuryak/vk-chatbot/internal/entities"
)

type (
	UsersRepo interface {
		Create(ctx context.Context, u entities.User) (*entities.User, error)
		GetByVKID(ctx context.Context, VKID int) (*entities.User, error)
		// TODO: Full CRUD
	}

	SympathyRepo interface {
		Create(ctx context.Context, s entities.Sympathy) (*entities.Sympathy, error)
		GetByUserIDs(ctx context.Context, firstUserID, secondUserID int) (*entities.Sympathy, error)
		UpdateReciprocity(ctx context.Context, id int, reciprocity bool) (*entities.Sympathy, error)
		// TODO: Full CRUD
	}
)
