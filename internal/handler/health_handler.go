package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func New() *Handler { return &Handler{} }

func (h *Handler) Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
}
