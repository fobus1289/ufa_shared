package http

type numberOrStr interface {
	int | int32 | int64 | int16 | int8 |
		uint | uint32 | uint64 | uint8 |
		string
}

type HttpResponseID[T numberOrStr] struct {
	Id T `json:"id"`
}

type HttpErrorResponse[T any] struct {
	Message T `json:"message"`
}

func ErrorResponse[T any](message T) *HttpErrorResponse[T] {
	return &HttpErrorResponse[T]{message}
}

func ID[T numberOrStr](id T) *HttpResponseID[T] {
	return &HttpResponseID[T]{id}
}
