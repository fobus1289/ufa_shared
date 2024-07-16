package service

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	"context"
	"github.com/fobus1289/ufa_shared/http/response"
	"gorm.io/gorm"
)

type ServiceScope = func(d *gorm.DB) *gorm.DB

type {{ $serviceNameUcWithService }} interface {
	FindOne(ctx context.Context, scopes ...ServiceScope) (*model.{{ $serviceNameUc }}Model, error)
	Find(ctx context.Context, scopes ...ServiceScope) ([]model.{{ $serviceNameUc }}Model, error)
	Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.Page{{ $serviceNameUc }}ResponseType, error)
	Create({{ $serviceNameLc }}Dto *dto.Create{{ $serviceNameUc }}Dto) (*response.ID, error)
	Update({{ $serviceNameLc }}Dto *dto.Update{{ $serviceNameUc }}Dto, scopes ...ServiceScope) error
	Delete(scopes ...ServiceScope) error
}

type {{ $serviceNameLcWithService }} struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) {{ $serviceNameUcWithService }} {
	return &{{ $serviceNameLcWithService }}{db}
}

func (s *{{ $serviceNameLcWithService }}) ModelWithContext(ctx context.Context) *gorm.DB {
	return s.db.WithContext(ctx).Model(&model.{{ $serviceNameUc }}Model{})
}

func (s *{{ $serviceNameLcWithService }}) Model() *gorm.DB {
	return s.db.Model(&model.{{ $serviceNameUc }}Model{})
}

func (s *{{ $serviceNameLcWithService }}) FindOne(ctx context.Context, scopes ...ServiceScope) (*model.{{ $serviceNameUc }}Model, error) {

	var {{ $serviceNameLc }} model.{{ $serviceNameUc }}Model
    {
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			First(&{{ $serviceNameLc }}).
			Error

		if err != nil {
			return nil, err
		}
    }

	return &{{ $serviceNameLc }}, nil
}

func (s *{{ $serviceNameLcWithService }}) Find(ctx context.Context, scopes ...ServiceScope) ([]model.{{ $serviceNameUc }}Model, error) {

	var moreModels []model.{{ $serviceNameUc }}Model
	{
		err := s.ModelWithContext(ctx).
			Scopes(scopes...).
			Find(&moreModels).
			Error

		if err != nil {
			return nil, err
		}
	}

	return moreModels, nil
}

func (s *{{ $serviceNameLcWithService }}) Page(ctx context.Context, take int, filter, limitFilter ServiceScope) (*dto.Page{{ $serviceNameUc }}ResponseType, error) {

	tx := s.ModelWithContext(ctx)

	var total int64
	{
		txTotal := tx.Scopes(filter).Count(&total)
		if err := txTotal.Error; err != nil {
			return nil, err
		}
	}

    var {{ $serviceNameLc }}s []*model.{{ $serviceNameUc }}Model
	{
		if err := tx.Scopes(filter, limitFilter).
			Find(&{{ $serviceNameLc }}s).Error; err != nil {
			return nil, err
		}
	}

	totalPages := int64(math.Ceil(float64(total) / float64(take)))

	return response.NewPaginateResponse(totalPages, {{ $serviceNameLc }}s), nil
}

func (s *{{ $serviceNameLcWithService }}) Create({{ $serviceNameLc }}Dto *dto.Create{{ $serviceNameUc }}Dto) (*response.ID, error) {

    {{ $serviceNameLc }} := model.{{ $serviceNameUc }}Model{
        Name: {{ $serviceNameLc }}Dto.Name,
    }

	if err := s.db.Create(&{{ $serviceNameLc }}).Error; err != nil {
		return nil, err
	}

	return &response.ID{Id: {{ $serviceNameLc }}.Id}, nil
}

func (s *{{ $serviceNameLcWithService }}) Update({{ $serviceNameLc }}Dto *dto.Update{{ $serviceNameUc }}Dto, scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Updates({{ $serviceNameLc }}Dto).Error
}

func (s *{{ $serviceNameLcWithService }}) Delete(scopes ...ServiceScope) error {
	return s.Model().Scopes(scopes...).Delete(nil).Error
}
