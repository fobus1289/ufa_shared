package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type Paginate struct{ page, perpage int }

func NewPaginate(page, perPage string) *Paginate {
	return &Paginate{
		page:    transformPage(page),
		perpage: transformPerPage(perPage),
	}
}

func NewPaginateEchoWithContext(ctx echo.Context) *Paginate {
	return &Paginate{
		page:    transformPage(ctx.QueryParam("page")),
		perpage: transformPerPage(ctx.QueryParam("perpage")),
	}
}

func (p *Paginate) Page() int {
	return p.page
}

func (p *Paginate) PerPage() int {
	return p.perpage
}

func (p *Paginate) Take() int {
	return p.perpage
}

func (p *Paginate) Skip() int {
	return p.perpage * p.page
}

func transformPerPage(perPage string) int {
	value, err := strconv.ParseInt(perPage, 10, 32)
	{
		if err != nil || value < 1 {
			value = 5
		} else if value > 25 {
			value = 25
		}
	}

	return int(value)
}

func transformPage(page string) int {
	value, err := strconv.ParseInt(page, 10, 32)
	{
		if err != nil || value < 0 {
			value = 0
		}
	}

	return int(value)
}
