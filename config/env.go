package config

import "sync"

var (
	envProject EnvProject
	envOs      EnvOs
)

type EnvProject struct {
	Mode        string `env:"MODE"`
	HttpPort    int    `env:"HTTP_PORT"`
	StoragePath string `env:"STORAGE_PATH"`

	PgDB       string `env:"POSTGRES_DB"`
	PgHost     string `env:"POSTGRES_HOST"`
	PgPort     string `env:"POSTGRES_PORT"`
	PgUser     string `env:"POSTGRES_USER"`
	PgPassword string `env:"POSTGRES_PASSWORD"`

	RedisAddr string `env:"REDIS_ADDR"`

	JwtSecret  string `env:"JWT_SECRET"`
	JwtExpire  int64  `env:"JWT_EXPIRE"`
	JwtRefresh int64  `env:"JWT_REFRESH"`

	OneIdBaseUrl      string `env:"ONE_ID_BASE_URL"`
	OneIdRedirectUrl  string `env:"ONE_ID_REDIRECT_URL"`
	OneIdClientId     string `env:"ONE_ID_CLIENT_ID"`
	OneIdClientSecret string `env:"ONE_ID_CLIENT_SECRET"`
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
