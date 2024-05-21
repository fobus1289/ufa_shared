package service
{{ $serviceUc:=ucFirst .ServiceName }}
{{ $serviceLc:=lcFirst .ServiceName }}
{{ $service:= printf "%s%s" $serviceLc "_service" }}
{{ $serviceCreateDto:= printf "dto.Create%s%s" $serviceUc "Dto" }}
{{ $serviceUpdateDto:= printf "dto.Update%s%s" $serviceUc "Dto" }}
{{ $serviceModel := printf "model.%s%s" $serviceUc "Model" }}
{{ $serviceModelPaginate := printf "%s%s%s" "Service" $serviceUc "ModelPaginate" }}
{{ $serviceInterface := printf "%s%s" $serviceLc "Service" }}
import (
	"context"

	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type {{$serviceModelPaginate}} = response.PaginateResponse[*{{$serviceModel}}]
type ServiceScope = func(d *gorm.DB) *gorm.DB

type {{ucFirst $serviceInterface}} interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModel}}, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]{{$serviceModel}}, error)
	Page(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModelPaginate}}, error)
	Create({{$serviceLc}} *{{$serviceModel}}) (int64, error)
	Update({{$serviceLc}} *{{$serviceModel}}, scopes ...ServiceScope) error
	Delete({{$serviceLc}} *{{$serviceModel}}, scopes ...ServiceScope) error
}


type {{$serviceInterface}} struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) {{ucFirst $serviceInterface}} {
	return &{{$serviceInterface}}{db}
}

func (s *{{$serviceInterface}}) FindOne(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModel}}, error) {

	var {{ $serviceLc }} {{$serviceModel}}

	err := s.db.Model(&{{$serviceModel}}{}).
	    WithContext(ctx).
		Scopes(scopes...).
		First(&{{ $serviceLc }}).
		Error

	if err != nil {
		return nil, err
	}

	return &{{ $serviceLc }}, nil
}

func (s *{{$serviceInterface}}) Find(ctx context.Context, scopes ...ServiceScope) ([]{{$serviceModel}}, error) {
	var moreModels []{{$serviceModel}}

	err := s.db.Model(&{{$serviceModel}}{}).
	    WithContext(ctx).
		Scopes(scopes...).
		Find(&moreModels).
		Error

	if err != nil {
		return nil, err
	}

	return moreModels, nil
}

func (s *{{$serviceInterface}}) Page(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModelPaginate}}, error) {

	var (
	    paginate = &{{$serviceModelPaginate}}{}
	)

	if err := s.db.WithContext(ctx).Scopes(scopes...).Model(&{{$serviceModel}}{}).Count(&paginate.Total).Find(&paginate.Data).Error; err != nil {
		return paginate, err
	}

	return paginate, nil
}

func (s *{{$serviceInterface}}) Create({{$serviceLc}} *{{$serviceModel}}) (int64, error) {

	err := s.db.Create({{$serviceLc}}).Error

	if err != nil {
		return 0, err
	}

	return {{$serviceLc}}.Id, nil
}

func (s *{{$serviceInterface}}) Update({{$serviceLc}} *{{$serviceModel}}, scopes ...ServiceScope) error {
	return s.db.Model({{$serviceLc}}).
		Scopes(scopes...).
		Updates({{$serviceLc}}).
		Error
}

func (s *{{$serviceInterface}}) Delete({{$serviceLc}} *{{$serviceModel}}, scopes ...ServiceScope) error {
	return s.db.Model({{$serviceLc}}).
		Scopes(scopes...).
		Delete(&{{$serviceModel}}{}).
		Error
}
