package jwtService

type JwtConfig struct {
	Secret         string
	Expired        int64
	RefreshExpired int64
}

func (j *JwtConfig) SecretData() string {
	return j.Secret
}

func (j *JwtConfig) Exp() int64 {
	return j.Expired
}

func (j *JwtConfig) RefreshExp() int64 {
	return j.RefreshExpired
}
