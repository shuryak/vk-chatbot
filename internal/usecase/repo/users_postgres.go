package repo

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/shuryak/vk-chatbot/internal/entities"
	"github.com/shuryak/vk-chatbot/internal/usecase"
	"github.com/shuryak/vk-chatbot/pkg/postgres"
)

type UsersRepo struct {
	*postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{pg}
}

// Check for implementation
var _ usecase.UsersRepo = (*UsersRepo)(nil)

func (ur UsersRepo) Create(ctx context.Context, u entities.User) (*entities.User, error) {
	sql, args, err := ur.Builder.
		Insert("users").
		Columns("vk_id, photo_id, name, age, city, interested_in, activated").
		Values(u.VKID, u.PhotoID, u.Name, u.Age, u.City, u.InterestedIn, u.Activated).
		Suffix("RETURNING \"vk_id\", \"photo_id\", \"name\", \"age\", \"city\", \"interested_in\", \"activated\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, sql, args...)
	dbUser := entities.User{}
	if err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated); err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - row.Scan: %v", err)
	}

	return &dbUser, nil
}

func (ur UsersRepo) GetByVKID(ctx context.Context, VKID int) (*entities.User, error) {
	sql, args, err := ur.Builder.
		Select("vk_id", "photo_id", "name", "age", "city", "interested_in", "activated").
		From("users").
		Where(squirrel.Eq{"vk_id": VKID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByVKID - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, sql, args...)
	dbUser := entities.User{}
	err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated)
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByVKID - row.Scan: %v", err)
	}

	return &dbUser, nil
}

func (ur UsersRepo) Update(ctx context.Context, VKID int, columns usecase.Columns) (*entities.User, error) {
	sql, args, err := ur.Builder.
		Update("users").
		SetMap(columns).
		Where(squirrel.Eq{"vk_id": VKID}).
		Suffix("RETURNING \"vk_id\", \"photo_id\", \"name\", \"age\", \"city\", \"interested_in\", \"activated\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Update - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, sql, args...)
	dbUser := entities.User{}
	if err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated); err != nil {
		return nil, fmt.Errorf("UsersPostgres - Update - row.Scan: %v", err)
	}

	return &dbUser, nil
}
