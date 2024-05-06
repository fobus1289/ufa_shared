package service
{{ $serviceNameLc:=lcFirst .ServiceName }}
{{ $service:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceLc:=lcFirst $service }}
import (
	{{ $serviceNameLc }}Service "{{$serviceLc}}/{{ $serviceNameLc }}/service"
	{{ $serviceNameLc }}Handler "{{$serviceLc}}/{{ $serviceNameLc }}/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewService(routerGroup *echo.Group, db *gorm.DB) {
	{{ $serviceNameLc }}Handler.NewHandler(routerGroup, {{ $serviceNameLc }}Service.NewService(db))
}
