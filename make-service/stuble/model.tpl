package model

import (
	"gorm.io/gorm"
	"time"
)
{{ $service:= printf "%s%s" .ServiceName "Model" }}
{{ $serviceUc:=ucFirst $service }}
type {{$serviceUc}} struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreateAt  time.Time   `json:"create_at" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time   `json:"updated_at" gorm:"autoUpdateTime:true,default:null"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

//other name or remove this method
//func ({{$serviceUc}}) TableName() string {
//	return "{{$serviceUc}}"
//}
