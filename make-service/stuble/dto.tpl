package dto

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	"github.com/fobus1289/ufa_shared/http/response"
	"{{ .ModPath }}/{{ $serviceNameSc }}/model"
)

type Page{{ $serviceNameUc }}ResponseType = response.PaginateResponse[*model.{{ $serviceNameUc }}Model] // @name Page{{ $serviceNameUc }}ResponseType

type Create{{ $serviceNameUc }}Dto struct {
    Name string `json:"name"`
} //@name Create{{ $serviceNameUc }}Dto

type Update{{ $serviceNameUc }}Dto struct {
    Name *string `json:"name"`
} //@name Update{{ $serviceNameUc }}Dto

type {{ $serviceNameUc }}QueryParams struct {
    Name *string `query:"name"`
} //@name {{ $serviceNameUc }}QueryParams

/*
func (j *YourDtoName) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &j)
}

func (j YourDtoName) Value() (driver.Value, error) {
	return json.RawMessage(nil).MarshalJSON()
}
*/