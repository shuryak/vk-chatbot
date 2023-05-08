package usecase

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/models"
)

type UsersUseCase struct {
	repo UsersRepo
}

func NewUsersUseCase(repo UsersRepo) *UsersUseCase {
	return &UsersUseCase{repo}
}

// Check for implementation
var _ Users = (*UsersUseCase)(nil)

func (uc UsersUseCase) Create(ctx context.Context, u models.User) (*models.User, error) {
	e, err := uc.repo.GetByVKID(ctx, u.ID)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - Create - uc.repo.Create: %v", err)
	}

	if e == nil {
		e, err = uc.repo.Create(ctx, entities.User{
			VKID:         u.ID,
			PhotoID:      u.PhotoID,
			Name:         u.Name,
			Age:          u.Age,
			City:         u.City,
			InterestedIn: u.InterestedIn,
			Activated:    u.Activated,
		})
		if err != nil {
			return nil, fmt.Errorf("UsersUseCase - Create - uc.repo.Create: %v", err)
		}
	}

	res := models.User{
		ID:           e.VKID,
		PhotoID:      e.PhotoID,
		Name:         e.Name,
		Age:          e.Age,
		City:         e.City,
		InterestedIn: e.InterestedIn,
		Activated:    e.Activated,
	}

	return &res, nil
}

func (uc UsersUseCase) GetByID(ctx context.Context, ID int) (*models.User, error) {
	e, err := uc.repo.GetByVKID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetByID - uc.repo.GetByID: %v", err)
	}
	if e == nil {
		return nil, nil
	}

	res := models.User{
		ID:           e.VKID,
		PhotoID:      e.PhotoID,
		Name:         e.Name,
		Age:          e.Age,
		City:         e.City,
		InterestedIn: e.InterestedIn,
		Activated:    e.Activated,
	}

	return &res, nil
}

func (uc UsersUseCase) GetExceptOf(ctx context.Context, count, offset int, IDs ...int) ([]models.User, error) {
	e, err := uc.repo.GetExceptOf(ctx, count, offset, IDs...)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - GetExceptOf - uc.repo.GetExceptOf: %v", err)
	}
	if e == nil {
		return nil, nil
	}

	var res []models.User
	for _, v := range e {
		res = append(res, models.User{
			ID:           v.VKID,
			PhotoID:      v.PhotoID,
			Name:         v.Name,
			Age:          v.Age,
			City:         v.City,
			InterestedIn: v.InterestedIn,
			Activated:    v.Activated,
		})
	}

	return res, nil
}

func (uc UsersUseCase) Update(ctx context.Context, u models.User) (*models.User, error) {
	c := NewUpdateBuilder()
	if u.PhotoID != "" {
		c.PhotoURL(u.PhotoID)
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

	e, err := uc.repo.Update(ctx, u.ID, c.Columns)
	if err != nil {
		return nil, fmt.Errorf("UsersUseCase - Update - uc.repo.Update: %v", err)
	}
	if e == nil {
		return nil, nil
	}

	res := models.User{
		ID:           e.VKID,
		PhotoID:      e.PhotoID,
		Name:         e.Name,
		Age:          e.Age,
		City:         e.City,
		InterestedIn: e.InterestedIn,
		Activated:    e.Activated,
	}

	return &res, nil
}
