package handler

import (
	"database/sql"
	"errors"
	"goServerPractice/internal/service/auth"
	"goServerPractice/internal/service/user"
	"goServerPractice/internal/transport"
	"goServerPractice/internal/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Login(c echo.Context) error {
	var req transport.LoginRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request format")
	}

	if err := c.Validate(&req); err != nil {
		var validErr *validator.ValidationError
		if errors.As(err, &validErr) {
			return c.JSON(http.StatusBadRequest, transport.ErrorResponse{
				Error: transport.ErrorBody{
					Code:    "VALIDATION_ERROR",
					Message: "入力内容に誤りがあります",
					Details: validErr.Details,
				},
			})
		}
	}

	u, err := user.FindByEmail(h.db, ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusUnauthorized, transport.ErrorResponse{
				Error: transport.ErrorBody{
					Code:    "INVALID_CREDENTIALS",
					Message: "メールアドレスまたはパスワードが正しくありません。",
				},
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "database error")
	}
	if err := auth.CheckHashPassword(u.PasswordHash, req.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, transport.ErrorResponse{
			Error: transport.ErrorBody{
				Code:    "INVALID_CREDENTIALS",
				Message: "メールアドレスまたはパスワードが正しくありません。",
			},
		})
	}

	accessToken, err := auth.GenerateAccessToken(u.ID, u.Email, h.cfg.JWTSecret, h.cfg.JWTExpires)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	refreshToken, err := auth.GenerateRefreshToken(u.ID, h.cfg.JWTSecret, h.cfg.RefreshTokenExpires)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, transport.LoginResponse{
		User: transport.UserDTO{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
