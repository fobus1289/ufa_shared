package jwtService

import (
	"errors"
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
	ParseTokenWithExpTime(string) (TokenMetadata, error)
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

func (s *jwtService) ParseToken(tokenString string) (TokenMetadata, error) {
	token, err := s.parseToken(tokenString)
	if err != nil {
		return TokenMetadata{}, err
	}

	return token.Claims.(TokenMetadata), nil
}

func (s *jwtService) ParseTokenWithExpTime(tokenString string) (TokenMetadata, error) {
	token, err := s.parseToken(tokenString)
	if err != nil {
		return TokenMetadata{}, err
	}

	if !token.Valid {
		return TokenMetadata{}, errors.New("token is not valid")

	}

	return token.Claims.(TokenMetadata), nil
}

func (s *jwtService) parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *jwtService) GenerateNewTokens(payload Payload[IUser]) (string, error) {
	accessToken, err := s.generateNewAccessToken(payload)
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
				ExpiresAt: jwt.NewNumericDate(now.Add(time.Second * time.Duration(s.config.SecretKeyExpireMinutes))),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		},
	)

	tokenString, err := token.SignedString([]byte(s.config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
