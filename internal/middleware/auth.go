package middleware

import (
	"goServerPractice/internal/service/auth"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Authorizationヘッダーを取得
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(401, map[string]string{
					"error": "認証が必要です",
				})
			}

			// 2. Bearerプレフィックスをチェック
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(401, map[string]string{
					"error": "無効な認証形式です",
				})
			}

			// 3. JWTトークンを解析、検証
			claims, err := auth.ParseAccessToken(tokenString, jwtSecret)
			if err != nil {
				return c.JSON(401, map[string]string{
					"error": "無効なトークンです",
				})
			}

			// 4. ユーザーIDをコンテキストに保存
			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)

			return next(c)
		}
	}
}
