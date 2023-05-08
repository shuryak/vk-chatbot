package usecase

import (
	"context"
	"fmt"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/models"
)

type SympathyUseCase struct {
	repo SympathyRepo
}

func NewSympathyUseCase(repo SympathyRepo) *SympathyUseCase {
	return &SympathyUseCase{repo}
}

// Check for implementation
var _ Sympathy = (*SympathyUseCase)(nil)

func (uc *SympathyUseCase) Create(ctx context.Context, s models.Sympathy) (*models.Sympathy, error) {
	e, err := uc.repo.GetByUserIDs(ctx, s.FirstUserID, s.SecondUserID)
	if err != nil {
		return nil, fmt.Errorf("SympathyUseCase - Create - uc.repo.GetByUserIDs: %v", err)
	}
	if e == nil {
		e, err = uc.repo.Create(ctx, entities.Sympathy{
			FirstUserVKID:  s.FirstUserID,
			SecondUserVKID: s.SecondUserID,
			Reciprocity:    s.Reciprocity,
		})
		if err != nil {
			return nil, fmt.Errorf("SympathyUseCase - Create - uc.repo.Create: %v", err)
		}
	}

	res := models.Sympathy{
		ID:           e.ID,
		FirstUserID:  e.FirstUserVKID,
		SecondUserID: e.SecondUserVKID,
		Reciprocity:  e.Reciprocity,
	}

	return &res, nil
}

func (uc *SympathyUseCase) GetByUserIDs(ctx context.Context, firstUserID, secondUserID int) (*models.Sympathy, error) {
	e, err := uc.repo.GetByUserIDs(ctx, firstUserID, secondUserID)
	if err != nil {
		return nil, fmt.Errorf("SympathyUseCase - GetByUserIDs - uc.repo.GetByUserIDs: %v", err)
	}
	if e == nil {
		return nil, nil
	}

	res := models.Sympathy{
		ID:           e.ID,
		FirstUserID:  e.FirstUserVKID,
		SecondUserID: e.SecondUserVKID,
		Reciprocity:  e.Reciprocity,
	}

	return &res, nil
}

func (uc *SympathyUseCase) UpdateReciprocity(ctx context.Context, id int, reciprocity bool) (*models.Sympathy, error) {
	e, err := uc.repo.UpdateReciprocity(ctx, id, reciprocity)
	if err != nil {
		return nil, fmt.Errorf("SympathyUseCase - UpdateReciprocity - uc.repo.UpdateReciprocity: %v", err)
	}
	if e == nil {
		return nil, nil
	}

	res := models.Sympathy{
		ID:           e.ID,
		FirstUserID:  e.FirstUserVKID,
		SecondUserID: e.SecondUserVKID,
		Reciprocity:  e.Reciprocity,
	}

	return &res, nil
}
