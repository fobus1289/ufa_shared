package service
{{ $serviceUc:=ucFirst .ServiceName }}
{{ $serviceLc:=lcFirst .ServiceName }}
{{ $service:= printf "%s%s" $serviceLc "_service" }}
{{ $serviceCreateDto:= printf "dto.Create%s%s" $serviceUc "Dto" }}
{{ $serviceUpdateDto:= printf "dto.Update%s%s" $serviceUc "Dto" }}
{{ $serviceModel := printf "models.%s%s" $serviceUc "Model" }}
{{ $serviceModelPaginate := printf "%s%s%s" "Service" $serviceUc "ModelPaginate" }}
{{ $serviceInterface := printf "%s%s" $serviceLc "Service" }}
import (
	"context"
	"{{$service}}/dto"
	"{{$service}}/models"

	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type {{$serviceModelPaginate}} = response.PaginateResponse[*{{$serviceModel}}]
type ServiceScope = func(d *gorm.DB) *gorm.DB

type {{ucFirst $serviceInterface}} interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModel}}, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]{{$serviceModel}}, error)
	Page(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModelPaginate}}, error)
	Create(dto *{{$serviceCreateDto}}) (int64, error)
	Update(dto *{{$serviceUpdateDto}}, scopes ...ServiceScope) error
	Delete(scopes ...ServiceScope) error
}


type {{$serviceInterface}} struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) {{ucFirst $serviceInterface}} {
	return &{{$serviceInterface}}{db}
}

func (s *{{$serviceInterface}}) FindOne(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModel}}, error) {

	var model {{$serviceModel}}

	err := s.db.Model({{$serviceModel}}{}).
		Scopes(scopes...).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *{{$serviceInterface}}) Find(ctx context.Context, scopes ...ServiceScope) ([]{{$serviceModel}}, error) {
	var moreModels []{{$serviceModel}}

	err := s.db.Model({{$serviceModel}}{}).
		Scopes(scopes...).
		Find(&moreModels).Error

	if err != nil {
		return nil, err
	}

	return moreModels, nil
}

func (s *{{$serviceInterface}}) Page(ctx context.Context, scopes ...ServiceScope) (*{{$serviceModelPaginate}}, error) {

	var (
		total int64
		model []*{{$serviceModel}}
	)

	if err := s.db.Scopes(scopes...).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Scopes(scopes...).Find(&model).Error; err != nil {
		return nil, err
	}

	paginate := &{{$serviceModelPaginate}}{
		Total: int(total),
		Data:  model,
	}

	return paginate, nil
}

func (s *{{$serviceInterface}}) Create(dto *{{$serviceCreateDto}}) (int64, error) {

	var id int64

	err := s.db.Model({{$serviceModel}}{}).
		Select("id").
		Create(dto).
		Scan(&id).Error

	if err != nil {
		return 0, nil
	}

	return id, nil
}

func (s *{{$serviceInterface}}) Update(dto *{{$serviceUpdateDto}}, scopes ...ServiceScope) error {
	return s.db.Model({{$serviceModel}}{}).
		Scopes(scopes...).
		Updates(dto).Error
}

func (s *{{$serviceInterface}}) Delete(scopes ...ServiceScope) error {
	return s.db.Model({{$serviceModel}}{}).
		Scopes(scopes...).
		Delete(nil).Error
}
