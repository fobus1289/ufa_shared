package main
{{ $service:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceUc:=ucFirst $service }}
{{ $serviceLc:=lcFirst $service }}
import (
	"{{$serviceLc}}/service"
	"{{$serviceLc}}/transport/http"

	pkgConfig "github.com/fobus1289/ufa_shared/config"
	"github.com/fobus1289/ufa_shared/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type config struct {
	Addr string `env:"SERVICE_PORT"`
}

type dbConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
	Port     uint   `env:"POSTGRES_PORT"`
}

func main() {

	dbConfigInstance := pkgConfig.Load(&dbConfig{})

	pgConfigInstance := pg.NewConfig(dbConfigInstance.Host,
		dbConfigInstance.User,
		dbConfigInstance.Password,
		dbConfigInstance.Database,
		dbConfigInstance.Port,
	)

	db, err := pg.NewGorm(pgConfigInstance)

	if err != nil {
		panic(err)
	}

	router := echo.New()

	setMiddlewares(router)

	group := router.Group("/api/v1")
	{
		http.NewHandler(group, service.NewService(db))
	}

	runHTTPServerOnAddr(router, pkgConfig.Load(&config{}).Addr)
}

func runHTTPServerOnAddr(handler *echo.Echo, port string) {

	if err := handler.Start(":" + port); err != nil {
		panic(err)
	}
}

func setMiddlewares(router *echo.Echo) {
	router.Use(middleware.RequestID())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
}
