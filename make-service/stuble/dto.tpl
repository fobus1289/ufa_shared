package dto
{{ $serviceLc:=lcFirst .ServiceName }}
{{ $service:= printf "%s%s" $serviceLc "_service" }}
{{ $serviceDto:= printf "%s%s" .ServiceName "Dto" }}
{{ $serviceDtoUc:=ucFirst $serviceDto }}
{{ $serviceUc:=ucFirst .ServiceName }}
{{ $serviceModel := printf "models.%s%s" $serviceUc "Model" }}

import "{{ $service }}/models"

type Create{{$serviceDtoUc}} struct {
	Name string `json:"name"`
}

func (c *Create{{ $serviceDtoUc }}) MarshalToDBModel() *{{ $serviceModel }} {
	return &{{ $serviceModel }}{
		Name:                 c.Name,
	}
}

type Update{{$serviceDtoUc}} struct {
	Name *string `json:"name"`
}
