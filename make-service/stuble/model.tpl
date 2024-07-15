package model

{{ $serviceNameSc := .ServiceName }}
{{ $serviceNameUc:=ucFirst .ServiceName }}
{{ $serviceNameLc:=lcFirst .ServiceName }}

{{ $serviceNameScWithService:= printf "%s%s" .ServiceName "_service" }}
{{ $serviceNameUcWithService:= ucFirst $serviceNameScWithService }}
{{ $serviceNameLcWithService:= lcFirst $serviceNameScWithService }}

import (
	"gorm.io/gorm"
	"time"
)

type {{ $serviceNameUc }}Model struct {
	Id        int64             `json:"id" gorm:"primaryKey"`
	Name      string            `json:"name" gorm:"unique"`
	CreatedAt *time.Time        `json:"createdAt" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time        `json:"updatedAt,omitempty" gorm:"autoUpdateTime:true"`
	DeletedAt *gorm.DeletedAt   `json:"-" swaggerignore:"true"`
}

func ({{ $serviceNameUc }}Model) TableName() string {
	return "{{ $serviceNameSc }}s"
}
