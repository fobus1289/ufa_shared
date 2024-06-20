package jwtService

type JwtConfig struct {
	Secret         string
	Expired        int64
	RefreshExpired int64
}
