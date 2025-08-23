package router

import (
	"goServerPractice/internal/config"
	"goServerPractice/internal/handler"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, cfg config.Config, h *handler.Handler) {
	// ヘルスチェック
	e.GET("/health", h.HealthCheck)
	e.POST("/users", h.CreateUser)
	e.POST("/login", h.Login)
}
