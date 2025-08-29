package repository

import (
	"context"
	"goServerPractice/internal/model"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserFilter struct {
	Status     *string
	Gender     *string
	Prefecture *string
	Limit      int
	Offset     int
	OrderBy    string
}

func (r *UserRepository) GetUsers(ctx context.Context, filter UserFilter) ([]model.User, int, error) {
	var users []model.User

	query := r.db.NewSelect().Model(&users)

	if filter.Status != nil && *filter.Status != "" {
		query = query.Where("status = ?", *filter.Status)
	}

	if filter.Gender != nil && *filter.Gender != "" {
		query = query.Where("gender = ?", *filter.Gender)
	}

	if filter.Prefecture != nil && *filter.Prefecture != "" {
		query = query.Where("prefecture = ?", *filter.Prefecture)
	}

	totalCount, err := query.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if filter.OrderBy == "" || filter.OrderBy == "updated_at_desc" {
		query = query.Order("updated_at DESC")
	}

	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	} else {
		query = query.Limit(20)
	}

	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err = query.Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}
