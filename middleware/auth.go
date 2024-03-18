package middleware

import "github.com/labstack/echo/v4"

type AccessToken string

type AuthService interface {
	VerifyToken(token AccessToken) error
}

type AuthMiddleware struct {
	authService AuthService
}

func NewAuthMiddleware(authService AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: authService}
}

func (m *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessToken := AccessToken(ctx.Request().Header.Get("Authorization"))
		if accessToken == "" {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		err := m.authService.VerifyToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(401, "Unauthorized")
		}
		return next(ctx)
	}
}
