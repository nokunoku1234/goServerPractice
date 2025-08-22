package handler

import (
	"errors"
	"goServerPractice/internal/service/auth"
	"goServerPractice/internal/service/user"
	"goServerPractice/internal/transport"
	"goServerPractice/internal/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	var req transport.CreateUserRequest
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

	emailExists, err := user.ExistsByEmail(h.db, ctx, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "database error")
	}
	if emailExists {
		return c.JSON(http.StatusConflict, transport.ErrorResponse{
			Error: transport.ErrorBody{
				Code:    "EMAIL_ALREADY_EXISTS",
				Message: "このメールアドレスは既に使用されています",
			},
		})
	}

	hashedPassword, err := auth.PasswordEncrypt(req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	u, err := h.db.User.
		Create().
		SetEmail(req.Email).
		SetPasswordHash(hashedPassword).
		SetName("名無し").
		Save(ctx)
	if err != nil {
		// 一意制約違反など（schemaでEmail.Unique()にしている想定）
		// ent.IsConstraintError(err) で分岐して 409 を返すなども可
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	accessToken, err := auth.GenerateAccessToken(u.ID, u.Email, h.cfg.JWTSecret, h.cfg.JWTExpires)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	refreshToken, err := auth.GenerateRefreshToken(u.ID, h.cfg.JWTSecret, h.cfg.RefreshTokenExpires)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, transport.CreateUserResponse{
		User: transport.UserDTO{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		},
		AccessToken:  accessToken,  // TODO: JWT実装後修正
		RefreshToken: refreshToken, // TODO: JWT実装後修正
	})
}
