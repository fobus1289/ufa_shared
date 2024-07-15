package dto

{{ $serviceNameSc := .ServiceName }}
{{ $serviceNameUc:=ucFirst .ServiceName }}
{{ $serviceNameLc:=lcFirst .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceNameUcWithService:= ucFirst $serviceNameScWithService }}
{{ $serviceNameLcWithService:= lcFirst $serviceNameScWithService }}

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
