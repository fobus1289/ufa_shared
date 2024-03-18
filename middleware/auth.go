package middleware

import (
	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	jwtService   jwtService.JwtService
	redisService redis.RedisService
}

const UnAuthorized = "Unauthorized"

func NewAuthMiddleware(jwtService jwtService.JwtService, redisService redis.RedisService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService:   jwtService,
		redisService: redisService,
	}
}

func (m *AuthMiddleware) Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		accessToken := ctx.Request().Header.Get(UnAuthorized)
		if accessToken == "" {
			return echo.NewHTTPError(401, UnAuthorized)
		}
		_, err := m.jwtService.ParseToken(accessToken)
		if err != nil {
			return echo.NewHTTPError(401, UnAuthorized)
		}
		return next(ctx)
	}
}
