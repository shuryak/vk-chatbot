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

type UsersRepo struct {
	*postgres.Postgres
}

func NewUsersRepo(pg *postgres.Postgres) *UsersRepo {
	return &UsersRepo{pg}
}

// Check for implementation
var _ usecase.UsersRepo = (*UsersRepo)(nil)

func (ur UsersRepo) Create(ctx context.Context, u entities.User) (*entities.User, error) {
	statement, args, err := ur.Builder.
		Insert("users").
		Columns("vk_id, photo_id, name, age, city, interested_in, activated").
		Values(u.VKID, u.PhotoID, u.Name, u.Age, u.City, u.InterestedIn, u.Activated).
		Suffix("RETURNING \"vk_id\", \"photo_id\", \"name\", \"age\", \"city\", \"interested_in\", \"activated\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, statement, args...)
	dbUser := entities.User{}
	err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Create - row.Scan: %v", err)
	}

	return &dbUser, nil
}

func (ur UsersRepo) GetByVKID(ctx context.Context, VKID int) (*entities.User, error) {
	statement, args, err := ur.Builder.
		Select("vk_id", "photo_id", "name", "age", "city", "interested_in", "activated").
		From("users").
		Where(squirrel.Eq{"vk_id": VKID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByID - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, statement, args...)
	dbUser := entities.User{}
	err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetByID - row.Scan: %v", err)
	}

	return &dbUser, nil
}

func (ur UsersRepo) GetExceptOf(ctx context.Context, count, offset int, IDs ...int) ([]entities.User, error) {
	statement, args, err := ur.Builder.
		Select("vk_id", "photo_id", "name", "age", "city", "interested_in", "activated").
		From("users").
		Where(squirrel.NotEq{"vk_id": IDs}).
		Offset(uint64(offset)).
		Limit(uint64(count)).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetExceptOf - row.Scan: %v", err)
	}

	rows, err := ur.Pool.Query(ctx, statement, args...)
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - GetExceptOf - ur.Pool.Query: %v", err)
	}
	defer rows.Close()

	dbUsers := make([]entities.User, 0, count)

	for rows.Next() {
		e := entities.User{}

		err = rows.Scan(&e.VKID, &e.PhotoID, &e.Name, &e.Age, &e.City, &e.InterestedIn, &e.Activated)
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("UsersPostgres - GetExceptOf - ur.Pool.Query: %v", err)
		}

		dbUsers = append(dbUsers, e)
	}

	return dbUsers, nil
}

func (ur UsersRepo) Update(ctx context.Context, VKID int, columns usecase.Columns) (*entities.User, error) {
	statement, args, err := ur.Builder.
		Update("users").
		SetMap(columns).
		Where(squirrel.Eq{"vk_id": VKID}).
		Suffix("RETURNING \"vk_id\", \"photo_id\", \"name\", \"age\", \"city\", \"interested_in\", \"activated\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Update - ur.Builder: %v", err)
	}

	row := ur.Pool.QueryRow(ctx, statement, args...)
	dbUser := entities.User{}
	err = row.Scan(&dbUser.VKID, &dbUser.PhotoID, &dbUser.Name, &dbUser.Age, &dbUser.City, &dbUser.InterestedIn, &dbUser.Activated)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("UsersPostgres - Update - row.Scan: %v", err)
	}

	return &dbUser, nil
}
