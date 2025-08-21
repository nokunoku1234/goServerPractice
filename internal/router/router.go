package router

import (
	"net/http"

	"goServerPractice/internal/config"
	"goServerPractice/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, cfg config.Config) {
	// ヘルスチェック
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	// v1 API
	api := e.Group("/api/v1")
	h := handler.New() // 依存があればここで渡す
	api.GET("/ping", h.Ping)
}
