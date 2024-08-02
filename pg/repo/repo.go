package repo

import (
	"fmt"

	"github.com/fobus1289/ufa_shared/http"
	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type where = func(tx *gorm.DB) *gorm.DB

func Page[T, M any](db *gorm.DB, paginate *http.Paginate, scopes ...where) (*response.PaginateResponse[T], error) {

	type PageEntity[T any] struct {
		Total int64
		Data  T `gorm:"embedded"`
	}

	var pageEntities []PageEntity[T]
	{
		scopes = append(scopes, func(tx *gorm.DB) *gorm.DB {
			tx.Statement.Selects = append(
				tx.Statement.Selects,
				fmt.Sprintf("CEIL(COUNT(1) OVER() / %f) AS total", float32(paginate.Take())),
			)
			return tx
		})

		var entityModel M

		if err := db.Model(entityModel).
			Scopes(scopes...).
			Offset(paginate.Skip()).
			Limit(paginate.Take()).
			Scan(&pageEntities).Error; err != nil {
			return nil, err
		}
	}

	var totalPages int64
	{
		if len(pageEntities) > 0 {
			pageEntity := pageEntities[0]
			totalPages = pageEntity.Total
		}
	}

	var entities = make([]T, 0, len(pageEntities))
	{
		for _, pageEntity := range pageEntities {
			entities = append(entities, pageEntity.Data)
		}
	}

	return response.NewPaginateResponse(totalPages, entities), nil
}
