package user

import (
	"context"
	"goServerPractice/ent"
	"goServerPractice/ent/user"
)

func ExistsByEmail(client *ent.Client, ctx context.Context, email string) (bool, error) {
	exists, err := client.User.Query().Where(user.EmailEQ(email)).Exist(ctx)
	return exists, err
}

func FindByEmail(client *ent.Client, ctx context.Context, email string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	return u, err
}
