package config

import "sync"

var (
	envProject EnvProject
	envOs      EnvOs
)

type EnvProject struct {
	PgPassword       string `env:"POSTGRES_PASSWORD"`
	JwtRefresh       string `env:"POSTGRES_DB"`
	RedisAddr        string `env:"REDIS_ADDR"`
	PgHOST           string `env:"POSTGRES_HOST"`
	OneIdRedirectUrl string `env:"ONE_ID_LOGIN_REDIRECT_BASE_URL"`
	PgUSER           string `env:"POSTGRES_USER"`
	OneIdBaseUrl     string `env:"ONE_ID_BASE_URL"`
	JwtSecret        string `env:"POSTGRES_DB"`
	Mode             string `env:"MODE"`
	JwtExpired       string `env:"POSTGRES_DB"`
	PgDB             string `env:"POSTGRES_DB"`
	OneId            string `env:"ONE_ID_CLIENT_ID"`
	OneIdSecret      string `env:"ONE_ID_CLIENT_SECRET"`
	HttpPort         int    `env:"HTTP_PORT"`
	PgPORT           int    `env:"POSTGRES_PORT"`
}

type EnvOs struct {
	Path          string `env:"PATH"`
	Home          string `env:"HOME"`
	User          string `env:"USER"`
	Shell         string `env:"SHELL"`
	Logname       string `env:"LOGNAME"`
	Pwd           string `env:"PWD"`
	Lang          string `env:"LANG"`
	Tz            string `env:"TZ"`
	Editor        string `env:"EDITOR"`
	Term          string `env:"TERM"`
	Display       string `env:"DISPLAY"`
	LdLibraryPath string `env:"LD_LIBRARY_PATH"`
	CFlags        string `env:"CFLAGS"`
	LdFlags       string `env:"LDFLAGS"`
	TmpDir        string `env:"TMPDIR"`
	XdgConfigHome string `env:"XDG_CONFIG_HOME"`
	XdgDataHome   string `env:"XDG_DATA_HOME"`
	Mail          string `env:"MAIL"`
}

func init() {
	sync.OnceFunc(func() {
		MustLoad(&envProject)
		MustLoad(&envOs)
	})()
}

type MergeEnv struct {
	EnvOs
	EnvProject
}

func Env() MergeEnv {
	return MergeEnv{
		EnvProject: ProjectEnv(),
		EnvOs:      OsEnv(),
	}
}

func ProjectEnv() EnvProject {
	return envProject
}

func OsEnv() EnvOs {
	return envOs
}
