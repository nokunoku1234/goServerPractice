package app

import (
	"context"
	"goServerPractice/ent"
	"goServerPractice/internal/config"
	"goServerPractice/internal/handler"
	"goServerPractice/internal/router"
	"log"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func Run() error {
	cfg := config.Load() // ポートやDB URLを読む
	client, err := ent.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("failed opening connection: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	h := handler.New(client)
	router.Register(e, cfg, h) // ルート登録
	return e.Start(":" + cfg.Port)
}
