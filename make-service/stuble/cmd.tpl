package main

{{ $serviceNameSc := .ServiceName }}
{{ $serviceNameUc:=ucFirst .ServiceName }}
{{ $serviceNameLc:=lcFirst .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceNameUcWithService:= ucFirst $serviceNameScWithService }}
{{ $serviceNameLcWithService:= lcFirst $serviceNameScWithService }}

import (
    "gorm.io/gorm"
	{{ $serviceNameLc }}Service "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/service"
    {{ $serviceNameLc }}Model "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/model"
    {{ $serviceNameLc }}Handler "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/handler"

	loader "github.com/fobus1289/ufa_shared/config"
	"github.com/fobus1289/ufa_shared/pg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	projectEnv := loader.ProjectEnv()

	pgConfig := pg.NewConfigEmpty()
	{
		pgConfig.SetHost(projectEnv.PgHost).
			SetPort(projectEnv.PgPort).
			SetDbname(projectEnv.PgDB).
			SetUser(projectEnv.PgUser).
			SetPassword(projectEnv.PgPassword)
	}

	db, err := pg.NewGorm(&pgConfig)
	{
        if err != nil {
            log.Panicln(err)
        }

        db.AutoMigrate(
            {{ $serviceNameLc }}Model.{{ $serviceNameUc }}Model{},
        )
	}

	router := echo.New()
	{
	    setMiddlewares(router)
	    createHandler(router, db)
	    runHTTPServerOnAddr(router, projectEnv.HttpPort)
	}
}

func runHTTPServerOnAddr(handler *echo.Echo, port int) {
	url := strconv.FormatInt(int64(port), 10)
	{
		log.Panicln(handler.Start(":" + url))
	}
}

func setMiddlewares(router *echo.Echo) {
	router.Use(middleware.RemoveTrailingSlash())
	router.Use(middleware.RequestID())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
}

func createHandler(router *echo.Echo, db *gorm.DB) {
	group := router.Group("/api/v1")
	{
		{{ $serviceNameLc }}Handler.NewHandler(group, {{ $serviceNameLc }}Service.NewService(db))
	}
}
