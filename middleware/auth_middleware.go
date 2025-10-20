package middleware

import (
	"errors"
	"strings"

	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware[U jwtService.IUser[T, E, K], T, E, K any] struct {
	jwtService   jwtService.JwtService[U, T, E, K]
	redisService redis.RedisService
	db           *gorm.DB
}

func NewAuthMiddleware[U jwtService.IUser[T, E, K], T, E, K any](
	jwtService jwtService.JwtService[U, T, E, K],
	redisService redis.RedisService,
	db *gorm.DB,
) *AuthMiddleware[U, T, E, K] {
	return &AuthMiddleware[U, T, E, K]{
		jwtService:   jwtService,
		redisService: redisService,
		db:           db,
	}
}

func AuthorizationToken(c echo.Context) (string, error) {
	const (
		UnAuthorizedErrorMessage = "unauthorized"
		Authorization            = "Authorization"
		Bearer                   = "Bearer "
	)

	token := c.Request().Header.Get(Authorization)
	{
		if token == "" || !strings.HasPrefix(token, Bearer) {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		if strings.Count(token, ".") != 2 {
			return "", errors.New(UnAuthorizedErrorMessage)
		}

		token = strings.TrimPrefix(token, Bearer)
	}

	return token, nil
}

