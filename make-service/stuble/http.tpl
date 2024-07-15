package service

{{ $serviceNameSc := .ServiceName }}
{{ $serviceNameUc:=ucFirst .ServiceName }}
{{ $serviceNameLc:=lcFirst .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceNameUcWithService:= ucFirst $serviceNameScWithService }}
{{ $serviceNameLcWithService:= lcFirst $serviceNameScWithService }}

import (
	{{ $serviceNameLc }}Service "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/service"
	{{ $serviceNameLc }}Handler "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewService(routerGroup *echo.Group, db *gorm.DB) {
	{{ $serviceNameLc }}Handler.NewHandler(routerGroup, {{ $serviceNameLc }}Service.NewService(db))
}
