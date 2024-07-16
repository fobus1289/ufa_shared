package dto

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	"github.com/fobus1289/ufa_shared/http/response"
	"{{ $serviceNameScWithService }}/{{ $serviceNameSc }}/model"
)

type Page{{ $serviceNameUc }}ResponseType = response.PaginateResponse[*model.{{ $serviceNameUc }}Model] // @name Page{{ $serviceNameUc }}ResponseType

type Create{{ $serviceNameUc }}Dto struct {
    Name string `json:"name"`
} //@name Create{{ $serviceNameUc }}Dto

type Update{{ $serviceNameUc }}Dto struct {
    Name *string `json:"name"`
} //@name Update{{ $serviceNameUc }}Dto
