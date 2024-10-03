package jwtService

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type IUser interface {
	ID() int64
	Pre(*gorm.DB, echo.Context) error
	PreWithPermission(*gorm.DB, echo.Context, ...string) error
}

type Payload[T IUser] struct {
	User T `json:"user"`
	jwt.RegisteredClaims
}

type JwtService[T IUser] interface {
	ParseToken(string) (T, error)
	ParseTokenWithExpired(string) (T, error)
	Token(T) (string, error)
	Config() JwtConfig
}

type jwtService[T IUser] struct {
	config JwtConfig
}

func (j *jwtService[T]) Token(user T) (string, error) {
	return Encode(user, j.config.Secret, j.config.Expired)
}

func (j *jwtService[T]) ParseToken(token string) (T, error) {
	payload, err := Decode[T](token, j.config.Secret, false)
	{
		if err != nil {
			var none T
			return none, err
		}
	}
	return payload.User, nil
}

func (j *jwtService[T]) ParseTokenWithExpired(token string) (T, error) {
	payload, err := Decode[T](token, j.config.Secret, true)
	{
		if err != nil {
			var none T
			return none, err
		}
	}
	return payload.User, nil
}

func (j *jwtService[T]) Config() JwtConfig {
	return j.config
}

func NewJwtService[T IUser](config JwtConfig) JwtService[T] {
	return &jwtService[T]{
		config: config,
	}
}

func Encode[T IUser](user T, secret string, expired int64) (string, error) {
	payload := NewPayload(user, expired)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString([]byte(secret))
	{
		if err != nil {
			return "", err
		}
	}

	return tokenString, nil
}

func Decode[T IUser](tokStr, secret string, withExpired bool) (*Payload[T], error) {
	keyfunc := func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	}

	token, err := jwt.ParseWithClaims(tokStr, &Payload[T]{}, keyfunc)
	{
		if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}

		if withExpired && !token.Valid {
			return nil, errors.New("token is not valid")
		}
	}

	payload, ok := token.Claims.(*Payload[T])
	{
		if !ok {
			return nil, errors.New("unknown error")
		}
	}

	return payload, nil
}

func NewPayload[T IUser](user T, expired int64) *Payload[T] {
	now := time.Now()

	exp := now.Add(time.Minute * time.Duration(expired))

	return &Payload[T]{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.FormatInt(user.ID(), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
}
