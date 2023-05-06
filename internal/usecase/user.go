package usecase

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/entities"
)

type UsersUseCase struct {
	repo UsersRepo
}

func NewUsersUseCase(repo UsersRepo) *UsersUseCase {
	return &UsersUseCase{repo}
}

// Check for implementation
var _ Users = (*UsersUseCase)(nil)

func (uc UsersUseCase) Create(ctx context.Context, u entities.User) (*entities.User, error) {
	e, err := uc.repo.Create(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - Create - uc.repo.Create: %v", err)
	}

	return e, nil
}
