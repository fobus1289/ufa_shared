package middleware

import (
	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/jwtService"
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/labstack/echo/v4"
)

// type AuthEchoMiddleware[U jwtService.IUser[T, E, K], T echo.Context, E, K any] interface {
// 	BuildMiddleware(permissions ...K) echo.MiddlewareFunc
// 	Redis() redis.RedisService
// 	Token(user U) (string, error)
// 	ParseToken(token string) (U, error)
// 	ParseTokenWithExpired(token string) (U, error)
// }

type AuthEchoMiddleware[U jwtService.IUser[T, E, K], T echo.Context, E, K any] struct {
	jwtService   jwtService.JwtService[U, T, E, K]
	redisService redis.RedisService
	storage      E
}

func NewAuthEchoMiddleware[U jwtService.IUser[T, E, K], T echo.Context, E, K any](
	config jwtService.JwtConfig,
	redisService redis.RedisService,
	storage E,
) *AuthEchoMiddleware[U, T, E, K] {

	jwtService := jwtService.NewJwtService[U](config)

	authEchoMiddleware := AuthEchoMiddleware[U, T, E, K]{
		jwtService:   jwtService,
		redisService: redisService,
		storage:      storage,
	}

	return &authEchoMiddleware
}

func (a *AuthEchoMiddleware[U, T, E, K]) BuildMiddleware(permissions ...K) echo.MiddlewareFunc {
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

			if err := user.Pre(ctx.(T), a.storage, permissions...); err != nil {
				return http.HTTPError(err).Unauthorized()
			}

			return next(ctx)
		}
	}
}

func (a *AuthEchoMiddleware[U, T, E, K]) Redis() redis.RedisService {
	return a.redisService
}

func (a *AuthEchoMiddleware[U, T, E, K]) Token(user U) (string, error) {
	return a.jwtService.Token(user)
}

func (a *AuthEchoMiddleware[U, T, E, K]) TokenWithExpired(user U, expired int64) (string, error) {
	secret := a.jwtService.Config().Secret
	return jwtService.Encode(user, secret, expired)
}

func (a *AuthEchoMiddleware[U, T, E, K]) ParseToken(token string) (U, error) {
	return a.jwtService.ParseToken(token)
}

func (a *AuthEchoMiddleware[U, T, E, K]) ParseTokenWithExpired(token string) (U, error) {
	return a.jwtService.ParseTokenWithExpired(token)
}

func (a *AuthEchoMiddleware[U, T, E, K]) ParseTokenWithGracePeriod(token string, grace int64) (U, error) {
	return a.jwtService.ParseTokenWithGracePeriod(token, grace)
}
