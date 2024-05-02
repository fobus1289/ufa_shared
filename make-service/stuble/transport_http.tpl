package http
{{ $serviceNameLc:=lcFirst .ServiceName }}
{{ $serviceNameUc:=ucFirst .ServiceName }}
{{ $service:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceLc:=lcFirst .ServiceName }}
{{ $handlerName:= printf "%s%s" $serviceNameLc "Handler" }}
{{ $serviceInterface := printf "%s%s" $serviceLc "Service" }}
{{ $serviceUc:=ucFirst $serviceNameUc }}
{{ $serviceCreateDto:= printf "dto.Create%s%s" $serviceUc "Dto" }}
{{ $serviceUpdateDto:= printf "dto.Update%s%s" $serviceUc "Dto" }}
{{ $serviceModel := printf "models.%s%s" $serviceUc "Model" }}
import (
	"{{$service}}/dto"
	"{{$service}}/service"

	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type {{$handlerName}} struct {
	service service.{{ucFirst $serviceInterface}}
}

func NewHandler(router *echo.Group, service service.{{ucFirst $serviceInterface}}) {

	group := router.Group("/{{$serviceNameLc}}")
	{
		handler := &{{$handlerName}}{service: service}

		group.POST("/", handler.Create)
		group.GET("/", handler.Page)
		group.GET("/:id", handler.GetById)
		group.PATCH("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

	}
}

// Create godoc
// @Summary      Create a new {{$serviceNameLc}}
// @Description  Create {{$serviceNameLc}}
// @Tags 		 {{$serviceNameLc}}
// @ID           create-{{$serviceNameLc}}
// @Accept       json
// @Produce      json
// @Param        input body {{$serviceCreateDto}} true "{{$serviceNameLc}} information"
// @Success      201 {object} int64 "Successful operation"
// @Failure      400 {object} map[string]interface{} "Bad request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /{{$serviceNameLc}} [post]
func (e *{{$handlerName}}) Create(ctx echo.Context) error {
	var createDto {{$serviceCreateDto}}

	if err := ctx.Bind(&createDto); err != nil {
		return http.Response(ctx).BadRequest(echo.Map{"message": err.Error()})
	}

	if err := validator.Validate(createDto); err != nil {
		return http.Response(ctx).BadRequest(
			echo.Map{"message": err.Error()},
		)
	}

	res, err := e.service.Create(createDto.MarshalToDBModel())
	if err != nil {
		return http.Response(ctx).InternalServerError(
			echo.Map{"error": err.Error()},
		)
	}

	return http.Response(ctx).Created(res)
}

// Page godoc
// @Summary      GetContent all {{$serviceNameLc}} with pagination
// @Description  GetContent all {{$serviceNameLc}} with pagination
// @Tags 		 {{$serviceNameLc}}
// @ID           get-all-{{$serviceNameLc}}
// @Accept       json
// @Produce      json
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Success      200 {object} map[string]interface{} "Successful operation"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /{{$serviceNameLc}} [get]
func (e *{{$handlerName}}) Page(ctx echo.Context) error {
	var (
		page     = ctx.QueryParam("page")
		perPage  = ctx.QueryParam("perpage")
		paginate = http.NewPaginate(page, perPage)
		rCtx     = ctx.Request().Context()
	)

	pageData, err := e.service.Page(rCtx, func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(paginate.Skip()).Limit(paginate.Take())
	})

	if err != nil {
		return http.Response(ctx).InternalServerError(
			echo.Map{"error": err.Error()},
		)
	}

	return http.Response(ctx).OK(pageData)
}

// GetById godoc
// @Summary      GetContent {{$serviceNameLc}} by ID
// @Description  GetContent {{$serviceNameLc}} by ID
// @Tags 		 {{$serviceNameLc}}
// @ID           get-{{$serviceNameLc}}-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "{{$serviceNameLc}} ID"
// @Success      200 {object} {{$serviceModel}} "Successful operation"
// @Failure      400,404 {object} map[string]interface{} "Bad request or Not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /{{$serviceNameLc}}/{id} [get]
func (e *{{$handlerName}}) GetById(ctx echo.Context) error {
	var (
		id         int64
		idStr      = ctx.Param("id")
		paramValue = http.PathValue(idStr)
	)

	if !paramValue.TryInt64(&id) {
		return http.Response(ctx).BadRequest(
			echo.Map{"message": "parse id error"},
		)
	}

	rCtx := ctx.Request().Context()

	organization, err := e.service.FindOne(rCtx, func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	})

	if err != nil {
		return http.Response(ctx).InternalServerError(
			echo.Map{"error": err.Error()},
		)
	}

	return http.Response(ctx).OK(organization)
}

// Update godoc
// @Summary      Update {{$serviceNameLc}} information
// @Description  Update {{$serviceNameLc}} information by ID
// @Tags 		 {{$serviceNameLc}}
// @ID           update-{{$serviceNameLc}}
// @Accept       json
// @Param        id path string true "{{$serviceNameLc}} ID"
// @Param        input body {{$serviceUpdateDto}} true "{{$serviceNameLc}} information"
// @Success      204 "Successful operation"
// @Failure      400 {object} map[string]interface{} "Bad request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /{{$serviceNameLc}}/{id} [patch]
func (e *{{$handlerName}}) Update(ctx echo.Context) error {
	var (
		id         int64
		idStr      = ctx.Param("id")
		paramValue = http.PathValue(idStr)
	)

	if !paramValue.TryInt64(&id) {
		return http.Response(ctx).BadRequest(
			echo.Map{"message": "parse id error"},
		)
	}

	var updateDto {{$serviceUpdateDto}}

	if err := ctx.Bind(&updateDto); err != nil {
		return http.Response(ctx).BadRequest(
			echo.Map{"message": err.Error()},
		)
	}

	err := e.service.Update(&updateDto, func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	})

	if err != nil {
		return http.Response(ctx).InternalServerError(
			echo.Map{"error": err.Error()},
		)
	}

	return http.Response(ctx).NoContent()
}

// Delete godoc
// @Summary      Delete {{$serviceNameLc}} by ID
// @Description  Delete {{$serviceNameLc}} by ID
// @Tags 		 {{$serviceNameLc}}
// @ID           delete-{{$serviceNameLc}}
// @Accept       json
// @Param        id path string true "{{$serviceNameLc}} ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} map[string]interface{} "Bad request"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /{{$serviceNameLc}}/{id} [delete]
func (e *{{$handlerName}}) Delete(ctx echo.Context) error {
	var (
		id         int64
		idStr      = ctx.Param("id")
		paramValue = http.PathValue(idStr)
	)

	if !paramValue.TryInt64(&id) {
		return http.Response(ctx).BadRequest(
			echo.Map{"message": "parse id error"},
		)
	}

	err := e.service.Delete(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	})

	if err != nil {
		return http.Response(ctx).InternalServerError(
			echo.Map{"error": err.Error()},
		)
	}

	return http.Response(ctx).NoContent()
}
