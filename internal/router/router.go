package router

import (
	"goServerPractice/internal/config"
	"goServerPractice/internal/handler"
	"goServerPractice/internal/middleware"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo, cfg config.Config, h *handler.Handler) {
	// ヘルスチェック
	e.GET("/health", h.HealthCheck)
	e.POST("/users", h.CreateUser)
	e.POST("/login", h.Login)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	protected.GET("/users/profile", h.GetUserProfile)
}
