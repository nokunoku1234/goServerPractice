package database

import (
	"context"
	"database/sql"
	"fmt"
	"goServerPractice/internal/model"
	"log"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDB(databaseURL string) (*bun.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	bunDb := bun.NewDB(db, pgdialect.New())
	bunDb.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true)))

	return bunDb, err
}

func RunMigration(db *bun.DB) error {
	ctx := context.Background()

	_, err := db.NewCreateTable().Model((*model.User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}
