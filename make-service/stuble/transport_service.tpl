package service
{{ $service:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceLc:=lcFirst $service }}
import (
	"{{$serviceLc}}/service"
	"{{$serviceLc}}/transport/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewService(routerGroup *echo.Group, db *gorm.DB) {
	http.NewHandler(routerGroup, service.NewService(db))
}
