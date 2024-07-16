package service

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	{{ $serviceNameLc }}Service "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/service"
	{{ $serviceNameLc }}Handler "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewService(routerGroup *echo.Group, db *gorm.DB) {
	{{ $serviceNameLc }}Handler.NewHandler(routerGroup, {{ $serviceNameLc }}Service.NewService(db))
}
