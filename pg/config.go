package pg

import "fmt"

type Config struct {
	host     string
	user     string
	password string
	dbname   string
	port     uint
	sslmode  bool
}

// Default
// host:     "postgres",
// user:     "postgres",
// password: "postgres",
// dbname:   "postgres",
// port:     5432,
// sslmode:  false,
func Default() Config {
	return Config{
		host:     "postgres",
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		port:     5432,
		sslmode:  false,
	}
}

func (c Config) ssl() string {
	if !c.sslmode {
		return "sslmode=disable"
	}
	return ""
}

func (c Config) build() string {
	return fmt.Sprintf(
		`host=%s
     port=%d 
     user=%s 
     password=%s 
     dbname=%s 
     %s`,
		c.host,
		c.port,
		c.user,
		c.password,
		c.dbname,
		c.ssl(),
	)
}

func (c *Config) SetHost(host string) *Config {
	c.host = host
	return c
}

func (c *Config) SetPort(port uint) *Config {
	c.port = port
	return c
}

func (c *Config) SetUser(user string) *Config {
	c.user = user
	return c
}

func (c *Config) SetPassword(password string) *Config {
	c.password = password
	return c
}

func (c *Config) SetDbname(dbname string) *Config {
	c.dbname = dbname
	return c
}

func (c *Config) SetSSL(sslmode bool) *Config {
	c.sslmode = sslmode
	return c
}

func NewConfigEmpty() Config {
	return Config{}
}

func NewConfigWith(
	host, user, password, dbname string, port uint, sslmode bool,
) Config {
	return Config{
		host, user, password, dbname, port, sslmode,
	}
}

func NewConfig(
	host, user, password, dbname string, port uint,
) Config {
	return Config{
		host, user, password, dbname, port, false,
	}
}
