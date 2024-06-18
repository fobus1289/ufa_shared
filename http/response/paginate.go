package response

import "github.com/fobus1289/ufa_shared/http"

type ID http.ResponseID[int64] //@name ID

type PaginateResponse[T any] struct {
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}

func NewPaginateResponse[T any](total int64, data []T) *PaginateResponse[T] {
	return &PaginateResponse[T]{
		Total: total,
		Data:  data,
	}
}
