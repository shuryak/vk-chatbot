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

func (uc UsersUseCase) GetByVKID(ctx context.Context, VKID int) (*entities.User, error) {
	e, err := uc.repo.GetByVKID(ctx, VKID)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetByVKID - uc.repo.GetByVKID: %v", err)
	}

	return e, nil
}

func (uc UsersUseCase) Update(ctx context.Context, u entities.User) (*entities.User, error) {
	c := NewUpdateBuilder()
	if u.PhotoURL != "" {
		c.PhotoURL(u.PhotoURL)
	}
	if u.Name != "" {
		c.Name(u.Name)
	}
	if u.Age != 0 {
		c.Age(u.Age)
	}
	if u.City != "" {
		c.City(u.City)
	}
	if u.InterestedIn != "" {
		c.InterestedIn(u.InterestedIn)
	}
	if u.Activated != nil {
		c.Activated(*u.Activated)
	}

	e, err := uc.repo.Update(ctx, u.VKID, c.Columns)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - Update - uc.repo.Update: %v", err)
	}

	return e, nil
}
