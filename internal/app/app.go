package app

import (
	"github.com/labstack/echo/v4"
	"goServerPractice/internal/config"
	"goServerPractice/internal/router"
)

func Run() error {
	cfg := config.Load() // ポートやDB URLを読む
	e := echo.New()
	router.Register(e, cfg) // ルート登録
	return e.Start(":" + cfg.Port)
}
