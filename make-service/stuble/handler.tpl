package handler

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}
{{ $serviceNameKc:=toKebab .ServiceName }}
{{ $serviceNameWithSpace:=withSpace .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	_ "github.com/fobus1289/ufa_shared/http/response"
	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	_ "{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/model"
)

type {{ $serviceNameLc }}Handler struct {
	service service.{{ $serviceNameUc }}Service
}

func NewHandler(router *echo.Group, service service.{{ $serviceNameUc }}Service) {

	group := router.Group("/{{ $serviceNameSc }}")
	{
		handler := &{{ $serviceNameLc }}Handler{service: service}

		group.POST("", handler.Create)
		group.GET("/page", handler.Page)
		group.GET("/:id", handler.GetById)
		group.PATCH("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

	}
}

// Create godoc
// @Summary      Create a new {{ $serviceNameWithSpace }}
// @Description  Create {{ $serviceNameWithSpace }}
// @Tags 		 {{ $serviceNameSc }}
// @ID           create-{{ $serviceNameKc }}
// @Accept       json
// @Produce      json
// @Param        input body dto.Create{{ $serviceNameUc }}Dto true "{{ $serviceNameWithSpace }} information"
// @Success      201 {object} response.ID "Successful operation"
// @Failure      400 {object} response.ErrorResponse "Bad request"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /{{ $serviceNameSc }} [post]
func (e *{{ $serviceNameLc }}Handler) Create(c echo.Context) error {
	var createDto dto.Create{{ $serviceNameUc }}Dto
	{
		if err := c.Bind(&createDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}

		if err := validator.Validate(createDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	idDto, err := e.service.Create(&createDto)
	{
		if err != nil {
			return http.HTTPError(err).InternalServerError()
		}
	}

	return http.Response(c).Created(idDto)
}

// Page godoc
// @Summary      GetContent all {{ $serviceNameWithSpace }} with pagination
// @Description  GetContent all {{ $serviceNameWithSpace }} with pagination
// @Tags 		 {{ $serviceNameSc }}
// @ID           get-all-{{ $serviceNameKc }}
// @Accept       json
// @Produce      json
// @Param        page query string false "Page number" default(1)
// @Param        perpage query string false "Number of items per page" default(10)
// @Param        search query string false "Searching by name or description"
// @Success      201 {object} response.ID "Successful operation"
// @Failure      400 {object} response.ErrorResponse "Bad request"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /{{ $serviceNameSc }}/page [get]
func (e *{{ $serviceNameLc }}Handler) Page(c echo.Context) error {
	var (
	    search   = c.QueryParam("search")
		page     = c.QueryParam("page")
		perPage  = c.QueryParam("perpage")
		paginate = http.NewPaginate(page, perPage)
		ctx     = c.Request().Context()
	)

	limitFilter := func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(paginate.Skip()).Limit(paginate.Take()).Order("id ASC")
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		if search != "" {
			search = fmt.Sprintf("%%%s%%", search)
			tx = tx.Where("name ILIKE ?", search)
		}
		return tx
	}

	pageData, err := e.service.Page(ctx, paginate.Take(), filter, limitFilter)
	{
		if err != nil {
			return http.HTTPError(err).InternalServerError()
		}
	}
	return http.Response(c).OK(pageData)
}

// GetById godoc
// @Summary      GetContent {{ $serviceNameWithSpace }} by ID
// @Description  GetContent {{ $serviceNameWithSpace }} by ID
// @Tags 		 {{ $serviceNameSc }}
// @ID           get-{{ $serviceNameKc }}-by-id
// @Accept       json
// @Produce      json
// @Param        id path string true "{{ $serviceNameWithSpace }} ID"
// @Success      200 {object} model.{{ $serviceNameUc }}Model "Successful operation"
// @Failure      400 {object} response.ErrorResponse "Bad request"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /{{ $serviceNameSc }}/{id} [get]
func (e *{{ $serviceNameLc }}Handler) GetById(c echo.Context) error {

	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("error parse id")
			return http.HTTPError(err).BadRequest()
		}
	}

	ctx := c.Request().Context()

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	}

	{{ $serviceNameLc }}, err := e.service.FindOne(ctx, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).OK({{ $serviceNameLc }})
}

// Update godoc
// @Summary      Update {{ $serviceNameWithSpace }} information
// @Description  Update {{ $serviceNameWithSpace }} information by ID
// @Tags 		 {{ $serviceNameSc }}
// @ID           update-{{ $serviceNameKc }}
// @Accept       json
// @Param        id path string true "{{ $serviceNameWithSpace }} ID"
// @Param        input body dto.Update{{ $serviceNameUc }}Dto true "{{ $serviceNameWithSpace }} information"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.ErrorResponse "Bad request"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /{{ $serviceNameSc }}/{id} [patch]
func (e *{{ $serviceNameLc }}Handler) Update(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	var updateDto dto.Update{{ $serviceNameUc }}Dto
	{
		if err := c.Bind(&updateDto); err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	filter := func(tx *gorm.DB) *gorm.DB {
		return tx.Where("id", id)
	}

	err := e.service.Update(&updateDto, filter)
	{
		if err != nil {
			return http.HTTPError(err).BadRequest()
		}
	}

	return http.Response(c).NoContent()
}

// Delete godoc
// @Summary      Delete {{ $serviceNameWithSpace }}
// @Description  Delete {{ $serviceNameWithSpace }} by ID
// @Tags 		 {{ $serviceNameSc }}
// @ID           delete-{{ $serviceNameKc }}
// @Accept       json
// @Param        id path string true "{{ $serviceNameWithSpace }} ID"
// @Success      204 "Successful operation"
// @Failure      400 {object} response.ErrorResponse "Bad request"
// @Failure      500 {object} response.ErrorResponse "Internal server error"
// @Router       /{{ $serviceNameSc }}/{id} [delete]
func (e *{{ $serviceNameLc }}Handler) Delete(c echo.Context) error {
	var id int64
	{
		if !http.PathValue(c.Param("id")).TryInt64(&id) {
			err := errors.New("parse id error")
			return http.HTTPError(err).BadRequest()
		}
	}

	{
		filter := func(tx *gorm.DB) *gorm.DB {
			return tx.Where("id", id)
		}

		if err := e.service.Delete(filter); err != nil {
			return http.HTTPError(err).InternalServerError()
		}
	}

	return http.Response(c).NoContent()
}
