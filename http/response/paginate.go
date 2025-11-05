package response

import "github.com/fobus1289/ufa_shared/http"

type ID http.HttpResponseID[int64] //@name ID

type ErrorResponse http.HttpErrorResponse[string] // @name ErrorResponse

type PaginateResponse[T any] struct {
	TotalCount int64 `json:"totalCount"`
	Total      int64 `json:"total"`
	Data       []T   `json:"data"`
}

func NewPaginateResponse[T any](totalCount, total int64, data []T) *PaginateResponse[T] {
	return &PaginateResponse[T]{
		TotalCount: totalCount,
		Total:      total,
		Data:       data,
	}
}
