package model

{{ $serviceNameSc :=toSnake .ServiceName }}
{{ $serviceNameUc:=toCamel .ServiceName }}
{{ $serviceNameLc:=toLowerCamel .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" $serviceNameSc "_service" }}
{{ $serviceNameUcWithService:= toCamel $serviceNameScWithService }}
{{ $serviceNameLcWithService:= toLowerCamel $serviceNameScWithService }}

import (
	"gorm.io/gorm"
	"time"
)

type {{ $serviceNameUc }}Model struct {
	Id        int64             `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name" gorm:"unique"`
	IsVisible bool             	`json:"isVisible" gorm:"default:true"`
	CreatedAt *time.Time        `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time        `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
	DeletedAt *gorm.DeletedAt   `json:"-" swaggerignore:"true"`
}

func ({{ $serviceNameUc }}Model) TableName() string {
	return "{{ $serviceNameSc }}s"
}
