package main

import (
	"goServerPractice/internal/config"
	"goServerPractice/internal/database"
	"goServerPractice/internal/seeder"
	"log"

	"github.com/uptrace/bun"
)

func main() {
	cfg := config.Load() // ポートやDB URLを読む
	db, err := database.NewDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
	}
	defer func(db *bun.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	if err := seeder.SeedUsers(db, 500); err != nil {
		log.Fatal(err)
	}

}
