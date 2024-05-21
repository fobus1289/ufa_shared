package dto
{{ $serviceLc:=lcFirst .ServiceName }}
{{ $service:= printf "%s%s" $serviceLc "_service" }}
{{ $serviceDto:= printf "%s%s" .ServiceName "Dto" }}
{{ $serviceDtoUc:=ucFirst $serviceDto }}
{{ $serviceUc:=ucFirst .ServiceName }}
{{ $serviceModel := printf "model.%s%s" $serviceUc "Model" }}



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

func (u *Update{{$serviceDtoUc}}) MarshalToDBModel({{ $serviceLc }} *{{ $serviceModel }}) {
    if u.Name != nil {
        {{ $serviceLc }}.Name = *u.Name
    }
}

type {{ $serviceUc }}ResponseDto struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt  time.Time   `json:"create_at"`
	UpdatedAt *time.Time   `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func (p *{{ $serviceUc }}ResponseDto) MarshalFromDbModel({{ $serviceLc }}Model {{ $serviceModel }}) {
    p.Id = {{ $serviceLc }}Model.Id
    p.Name = {{ $serviceLc }}Model.Name
    p.CreatedAt = {{ $serviceLc }}Model.CreatedAt
	p.UpdatedAt = {{ $serviceLc }}Model.UpdatedAt
	if {{ $serviceLc }}Model.DeletedAt.Valid {
		p.DeletedAt = &{{ $serviceLc }}Model.DeletedAt.Time
	} else {
		p.DeletedAt = nil
	}
}

func ConvertToPageDataResponse(pageData response.PaginateResponse[*{{ $serviceModel }}]) response.PaginateResponse[*{{ $serviceUc }}ResponseDto] {
    var dtoData []*{{ $serviceUc }}ResponseDto

    for _, {{ $serviceLc }}Model := range pageData.Data {
        {{ $serviceLc }}Dto := &{{ $serviceUc }}ResponseDto{}
        {{ $serviceLc }}Dto.MarshalFromDbModel(*{{ $serviceLc }}Model)
        dtoData = append(dtoData, {{ $serviceLc }}Dto)
    }

    return response.PaginateResponse[*{{ $serviceUc }}ResponseDto] {
        Total: pageData.Total,
        Data: dtoData,
    }
}
