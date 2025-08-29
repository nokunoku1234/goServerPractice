package handler

import (
	"goServerPractice/internal/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserList(c echo.Context) error {
	var req user.UserListRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request parameters",
		})
	}

	ctx := c.Request().Context()
	resp, err := user.GetUserList(h.db, ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch users",
		})
	}

	return c.JSON(http.StatusOK, resp)
}
