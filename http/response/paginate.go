package response

type PaginateResponse[T any] struct {
	Total int `json:"total"`
	Data  []T `json:"data"`
}
