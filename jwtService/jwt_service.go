package jwtService

import (
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtConfig struct {
	SecretKeyExpireSeconds  uint16 `env:"JWT_SECRET_KEY_EXPIRE_SECONDS"`
	SecretKey               string `env:"JWT_SECRET_KEY"`
	RefreshKeyExpireMinutes uint16 `env:"JWT_REFRESH_KEY_EXPIRE_MINUTES"`
	RefreshKey              string `env:"JWT_REFRESH_KEY"`
}

type TokenMetadata struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

type UserPayload struct {
	ID   string
	Role string
}

type JwtService interface {
	ExtractTokenMetadata(string) (TokenMetadata, error)
	GenerateNewTokens(UserPayload) (string, error)
}

type jwtService struct {
	config       *JwtConfig
	redisService redis.RedisService
}

func NewJwtService(config *JwtConfig) JwtService {
	return &jwtService{
		config: config,
	}
}

func (s *jwtService) ExtractTokenMetadata(token string) (TokenMetadata, error) {
	var claims TokenMetadata
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		return TokenMetadata{}, err
	}

	return claims, nil
}

func (s *jwtService) GenerateNewTokens(payload UserPayload) (string, error) {
	accessToken, err := s.generateNewAccessToken(payload)
	if err != nil {
		return "", err
	}

	err = s.redisService.SetWithTTL(payload.ID, accessToken, time.Minute*time.Duration(s.config.RefreshKeyExpireMinutes))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *jwtService) generateNewAccessToken(user UserPayload) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenMetadata{
			Role: user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        user.ID,
				Issuer:    s.config.SecretKey,
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(s.config.SecretKeyExpireSeconds))),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.Raw, nil
}

func (s *jwtService) generateNewRefreshToken(user UserPayload) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenMetadata{
			Role: user.Role,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        user.ID,
				Issuer:    s.config.SecretKey,
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * time.Duration(s.config.RefreshKeyExpireMinutes))),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.Raw, nil
}
