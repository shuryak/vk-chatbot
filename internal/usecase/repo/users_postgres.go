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
		Columns("vk_id, photo_url, name, city, interested_in").
		Values(u.VKID, u.PhotoURL, u.Name, u.City, u.InterestedIn).
		Suffix("RETURNING \"id\", \"vk_id\", \"photo_url\", \"name\", \"city\", \"interested_in\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, sql, args...)
	dbUser := entities.User{}
	if err = row.Scan(&dbUser.ID, &dbUser.VKID, &dbUser.PhotoURL, &dbUser.Name, &dbUser.City, &dbUser.InterestedIn); err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - row.Scan: %v", err)
	}

	return &dbUser, nil
}

func (ur UsersRepo) GetByVKID(ctx context.Context, VKID int) (*entities.User, error) {
	sql, args, err := ur.Builder.
		Select("id", "vk_id", "photo_url", "name", "city", "interested_in").
		From("users").
		Where(squirrel.Eq{"vk_id": VKID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByVKID - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, sql, args...)
	dbUser := entities.User{}
	err = row.Scan(&dbUser.ID, &dbUser.VKID, &dbUser.PhotoURL, &dbUser.Name, &dbUser.City, &dbUser.InterestedIn)
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByVKID - row.Scan: %v", err)
	}

	return &dbUser, nil
}
