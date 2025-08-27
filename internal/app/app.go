package app

import (
	"goServerPractice/internal/config"
	"goServerPractice/internal/database"
	"goServerPractice/internal/handler"
	"goServerPractice/internal/router"
	"goServerPractice/internal/validator"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"

	_ "github.com/lib/pq"
)

func Run() error {
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

	if err := database.RunMigration(db); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Validator = validator.NewCustomValidator()
	h := handler.New(db, cfg)
	router.Register(e, cfg, h) // ルート登録
	return e.Start(":" + cfg.Port)
}
