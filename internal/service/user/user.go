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
