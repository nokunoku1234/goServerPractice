package user

import (
	"context"
	"goServerPractice/internal/model"

	"github.com/uptrace/bun"
)

func ExistsByEmail(db *bun.DB, ctx context.Context, email string) (bool, error) {
	exists, err := db.NewSelect().Model((*model.User)(nil)).Where("email = ?", email).Exists(ctx)
	return exists, err
}

func FindByEmail(db *bun.DB, ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := db.NewSelect().Model(&user).Where("email = ?", email).Scan(ctx)
	return &user, err
}

func FindByID(db *bun.DB, ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
