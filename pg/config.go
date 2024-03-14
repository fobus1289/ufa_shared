package pg

import "fmt"

type connectionConfig struct {
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
func Default() connectionConfig {
	return connectionConfig{
		host:     "postgres",
		user:     "postgres",
		password: "postgres",
		dbname:   "postgres",
		port:     5432,
		sslmode:  false,
	}
}

func (c connectionConfig) ssl() string {
	if !c.sslmode {
		return "sslmode=disable"
	}
	return ""
}

func (c connectionConfig) build() string {
	return fmt.Sprintf(
		`host=%s
     port=%d 
     user=%s 
     password=%s 
     dbname=%s 
     %s
    `,
		c.host,
		c.port,
		c.user,
		c.password,
		c.dbname,
		c.ssl(),
	)
}

func (c *connectionConfig) SetHost(host string) *connectionConfig {
	c.host = host
	return c
}

func (c *connectionConfig) SetPort(port uint) *connectionConfig {
	c.port = port
	return c
}

func (c *connectionConfig) SetUser(user string) *connectionConfig {
	c.user = user
	return c
}

func (c *connectionConfig) SetPassword(password string) *connectionConfig {
	c.password = password
	return c
}

func (c *connectionConfig) SetDbname(dbname string) *connectionConfig {
	c.dbname = dbname
	return c
}

func (c *connectionConfig) SetSSL(sslmode bool) *connectionConfig {
	c.sslmode = sslmode
	return c
}

func NewConfigEmpty() connectionConfig {
	return connectionConfig{}
}

func NewConfigWith(
	host, user, password, dbname string, port uint, sslmode bool,
) connectionConfig {
	return connectionConfig{
		host, user, password, dbname, port, sslmode,
	}
}

func NewConfig(
	host, user, password, dbname string, port uint,
) connectionConfig {
	return connectionConfig{
		host, user, password, dbname, port, false,
	}
}
