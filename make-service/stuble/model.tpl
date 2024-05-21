package model

import (
	"gorm.io/gorm"
	"time"
)
{{ $serviceModel:= printf "%s%s" .ServiceName "Model" }}
{{ $serviceModelUc:=ucFirst $serviceModel }}
{{ $serviceLc:=lcFirst .ServiceName }}

type {{$serviceModelUc}} struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreatedAt  time.Time   `json:"create_at" gorm:"autoCreateTime:true"`
	UpdatedAt *time.Time   `json:"updated_at" gorm:"autoUpdateTime:true,default:null"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

//other name or remove this method
//func ({{$serviceModelUc}}) TableName() string {
//	return "{{$serviceModelUc}}"
//}

//func ({{ $serviceLc }} *{{ $serviceModelUc }}) BeforeCreate(tx *gorm.DB) (err error) {
//    return nil
//}

//func ({{ $serviceLc }} *{{ $serviceModelUc }}) BeforeUpdate(tx *gorm.DB) (err error) {
//    return nil
//}

//func ({{ $serviceLc }} *{{ $serviceModelUc }}) BeforeDelete(tx *gorm.DB) (err error) {
//    return nil
//}
