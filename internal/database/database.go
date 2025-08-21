package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"goServerPractice/ent"
	_ "github.com/lib/pq"
)

func NewClient(databaseURL string) (*ent.Client, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}
	
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	
	return client, nil
}

func RunMigration(client *ent.Client) error {
	ctx := context.Background()
	
	if err := client.Schema.Create(ctx); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}
	
	log.Println("Database migration completed successfully")
	return nil
}