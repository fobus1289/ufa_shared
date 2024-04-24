package dto
{{ $service:= printf "%s%s" .ServiceName "Dto" }}
{{ $serviceUc:=ucFirst $service }}
type Create{{$serviceUc}} struct {
	Name string `json:"name"`
}

type Update{{$serviceUc}} struct {
	Name *string `json:"name"`
}
