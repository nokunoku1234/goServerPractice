package seeder

import (
	"context"
	"fmt"
	"goServerPractice/internal/model"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

var genders = []string{"male", "female", "other"}
var statuses = []string{"active", "inactive", "suspended"}
var prefectures = []string{
	"北海道", "青森県", "岩手県", "宮城県", "秋田県",
	"山形県", "福島県", "茨城県", "栃木県", "群馬県",
	"埼玉県", "千葉県", "東京都", "神奈川県", "新潟県",
	"富山県", "石川県", "福井県", "山梨県", "長野県",
	"岐阜県", "静岡県", "愛知県", "三重県", "滋賀県",
	"京都府", "大阪府", "兵庫県", "奈良県", "和歌山県",
	"鳥取県", "島根県", "岡山県", "広島県", "山口県",
	"徳島県", "香川県", "愛媛県", "高知県", "福岡県",
	"佐賀県", "長崎県", "熊本県", "大分県", "宮崎県",
	"鹿児島県", "沖縄県",
}

func SeedUsers(db *bun.DB, count int) error {
	users := generateMockUsers(count)
	return insertUsers(db, users)
}

func generateMockUsers(count int) []model.User {
	users := make([]model.User, count)
	defaultPassword := "aaaaaaaa"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)

	for i := 0; i < count; i++ {
		daysAgo := rand.Intn(365)
		updatedAt := time.Now().AddDate(0, 0, -daysAgo)

		createdDaysAgo := daysAgo + rand.Intn(365)
		createdAt := time.Now().AddDate(0, 0, -createdDaysAgo)

		users[i] = model.User{
			Name:         gofakeit.Name(),
			Email:        fmt.Sprintf("user%d_%s", i+1, gofakeit.Email()),
			PasswordHash: string(hashedPassword),
			Bio:          gofakeit.Sentence(20),
			Status:       statuses[rand.Intn(len(statuses))],
			Gender:       genders[rand.Intn(len(genders))],
			Prefecture:   prefectures[rand.Intn(len(prefectures))],
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
		}
	}

	return users
}

func insertUsers(db *bun.DB, users []model.User) error {
	ctx := context.Background()
	batchSize := 100

	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		batch := users[i:end]
		_, err := db.NewInsert().Model(&batch).Exec(ctx)

		if err != nil {
			return fmt.Errorf("batch %d-%d failed: %w", i, end, err)
		}

		fmt.Printf("Inserted batch %d-%d\n", i+1, end)
	}

	return nil
}
