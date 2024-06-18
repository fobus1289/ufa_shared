package http

type numberOrStr interface {
	int | int32 | int64 | int16 | int8 |
		uint | uint32 | uint64 | uint8 |
		string
}

type ResponseID[T numberOrStr] struct {
	Id T `json:"id"`
}

type errorResponse[T any] struct {
	Message T `json:"message"`
}

func ErrorResponse[T any](message T) *errorResponse[T] {
	return &errorResponse[T]{message}
}

func ID[T numberOrStr](id T) *ResponseID[T] {
	return &ResponseID[T]{id}
}
