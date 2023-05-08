package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/postgres"
)

type SympathyRepo struct {
	*postgres.Postgres
}

func NewSympathyRepo(pg *postgres.Postgres) *SympathyRepo {
	return &SympathyRepo{pg}
}

// Check for implementation
var _ usecase.SympathyRepo = (*SympathyRepo)(nil)

func (sr SympathyRepo) Create(ctx context.Context, s entities.Sympathy) (*entities.Sympathy, error) {
	statement, args, err := sr.Builder.
		Insert("sympathy").
		Columns("first_user_vk_id", "second_user_vk_id", "reciprocity").
		Values(s.FirstUserVKID, s.SecondUserVKID, s.Reciprocity).
		Suffix("RETURNING \"id\", \"first_user_vk_id\", \"second_user_vk_id\", \"reciprocity\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - Create - sr.Builder: %v", err)
	}

	row := sr.Pool.QueryRow(ctx, statement, args...)
	dbSympathy := entities.Sympathy{}
	err = row.Scan(&dbSympathy.ID, &dbSympathy.FirstUserVKID, &dbSympathy.SecondUserVKID, &dbSympathy.Reciprocity)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - Create - row.Scan: %v", err)
	}

	return &dbSympathy, nil
}

func (sr SympathyRepo) GetByUserIDs(ctx context.Context, firstUserID, secondUserID int) (*entities.Sympathy, error) {
	statement, args, err := sr.Builder.
		Select("id", "first_user_vk_id", "second_user_vk_id", "reciprocity").
		From("sympathy").
		Where(squirrel.Or{
			squirrel.And{
				squirrel.Eq{"first_user_vk_id": firstUserID},
				squirrel.Eq{"second_user_vk_id": secondUserID},
			},
			squirrel.And{
				squirrel.Eq{"first_user_vk_id": secondUserID},
				squirrel.Eq{"second_user_vk_id": firstUserID},
			},
		}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - GetByUserIDs - ur.Builder: %v", err)
	}

	row := sr.Pool.QueryRow(ctx, statement, args...)
	dbSympathy := entities.Sympathy{}
	err = row.Scan(&dbSympathy.ID, &dbSympathy.FirstUserVKID, &dbSympathy.SecondUserVKID, &dbSympathy.Reciprocity)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - GetByUserIDs - row.Scan: %v", err)
	}

	return &dbSympathy, nil
}

func (sr SympathyRepo) UpdateReciprocity(ctx context.Context, id int, reciprocity bool) (*entities.Sympathy, error) {
	statement, args, err := sr.Builder.
		Update("sympathy").
		SetMap(map[string]interface{}{
			"reciprocity": reciprocity,
		}).
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING \"id\", \"first_user_vk_id\", \"second_user_vk_id\", \"reciprocity\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - UpdateReciprocity - sr.Builder: %v", err)
	}

	row := sr.Pool.QueryRow(ctx, statement, args...)
	dbSympathy := entities.Sympathy{}
	err = row.Scan(&dbSympathy.ID, &dbSympathy.FirstUserVKID, &dbSympathy.SecondUserVKID, &dbSympathy.Reciprocity)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("SympathyPostgres - UpdateReciprocity - row.Scan: %v", err)
	}

	return &dbSympathy, nil
}
