package jwtService

import (
	"github.com/fobus1289/ufa_shared/redis"
	"github.com/fobus1289/ufa_shared/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtConfig struct {
	SecretKeyExpireMinutes  uint16 `env:"JWT_SECRET_KEY_EXPIRE_MINUTES"`
	SecretKey               string `env:"JWT_SECRET_KEY"`
	RefreshKeyExpireMinutes uint16 `env:"JWT_REFRESH_KEY_EXPIRE_MINUTES"`
}

type IUser interface {
	ID() int64
}

type TokenMetadata struct {
	user IUser
	jwt.RegisteredClaims
}

type Payload[T IUser] struct {
	IUser
}

type JwtService interface {
	ParseToken(string) (TokenMetadata, error)
	GenerateNewTokens(Payload[IUser]) (string, error)
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

func (s *jwtService) ParseToken(token string) (TokenMetadata, error) {
	var claims TokenMetadata
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})

	if err != nil {
		return TokenMetadata{}, err
	}

	return claims, nil
}

func (s *jwtService) GenerateNewTokens(payload Payload[IUser]) (string, error) {
	accessToken, err := s.generateNewAccessToken(payload)
	if err != nil {
		return "", err
	}

	err = s.redisService.SetWithTTL(payload.ID(), accessToken, time.Minute*time.Duration(s.config.RefreshKeyExpireMinutes))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *jwtService) generateNewAccessToken(user Payload[IUser]) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenMetadata{
			user: user,
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        utils.Int64ToString(user.ID()),
				Issuer:    s.config.SecretKey,
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(s.config.SecretKeyExpireMinutes))),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.Raw, nil
}

func (s *jwtService) generateNewRefreshToken(user Payload[IUser]) (string, error) {
	now := time.Now()

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		TokenMetadata{
			RegisteredClaims: jwt.RegisteredClaims{
				ID:        utils.Int64ToString(user.ID()),
				Issuer:    s.config.SecretKey,
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Minute * time.Duration(s.config.RefreshKeyExpireMinutes))),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	return token.Raw, nil
}
