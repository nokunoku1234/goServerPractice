package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type createUserReq struct {
	Name string `json:"name"  validate:"required"`
}

type userRes struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req createUserReq
	if err := c.Bind(&req); err != nil || req.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body: name is required")
	}

	ctx := c.Request().Context()
	u, err := h.db.User.
		Create().
		SetName(req.Name).Save(ctx)
	if err != nil {
		// 一意制約違反など（schemaでEmail.Unique()にしている想定）
		// ent.IsConstraintError(err) で分岐して 409 を返すなども可
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	})
}
