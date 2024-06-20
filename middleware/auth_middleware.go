package middleware

import (
	"errors"
	"strings"

	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthMiddleware[T jwtService.IUser] struct {
	jwtService   jwtService.JwtService[T]
	redisService redis.RedisService
	db           *gorm.DB
}

func NewAuthMiddleware[T jwtService.IUser](
	jwtService jwtService.JwtService[T],
	redisService redis.RedisService,
	db *gorm.DB,
) *AuthMiddleware[T] {
	return &AuthMiddleware[T]{
		jwtService:   jwtService,
		redisService: redisService,
		db:           db,
	}
}

func (a *AuthMiddleware[T]) JwtPermissionMiddleware(permissions ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			token, err := AuthorizationToken(ctx)
			{
				if err != nil {
					return http.HTTPError(err).Unauthorized()
				}
			}

			user, err := a.jwtService.ParseTokenWithExpired(token)
			{
				if err != nil {
					return http.HTTPError(err).Unauthorized()
				}
			}

			if err := user.PreWithPermission(a.db, ctx, permissions...); err != nil {
				return http.HTTPError(err).Unauthorized()
			}

			return next(ctx)
		}
	}
}

func (a *AuthMiddleware[T]) JwtAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := AuthorizationToken(ctx)
		{
			if err != nil {
				return http.HTTPError(err).Unauthorized()
			}
			user, err := a.jwtService.ParseTokenWithExpired(token)
			{
				if err != nil {
					return http.HTTPError(err).Unauthorized()
				}
			}

			if err := user.Pre(a.db, ctx); err != nil {
				return http.HTTPError(err).Unauthorized()
			}
		}

		return next(ctx)
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
