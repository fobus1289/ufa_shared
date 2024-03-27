package response

type PaginateResponse[T any] struct {
	Total int `json:"total"`
	Data  []T `json:"data"`
}

func NewPaginateResponse[T any](total int, data []T) *PaginateResponse[T] {
	return &PaginateResponse[T]{
		Total: total,
		Data:  data,
	}
}
