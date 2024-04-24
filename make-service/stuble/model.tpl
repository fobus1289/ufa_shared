package models

import "gorm.io/gorm"
{{ $service:= printf "%s%s" .ServiceName "Model" }}
{{ $serviceUc:=ucFirst $service }}
type {{$serviceUc}} struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreateAt  bool   `json:"create_at" gorm:"autoCreateTime:true"`
	UpdatedAt bool   `json:"updated_at" gorm:"autoUpdateTime:true,default:null"`
	DeletedAt gorm.DeletedAt
}

//other name or remove this method
//func ({{$serviceUc}}) TableName() string {
//	return "{{$serviceUc}}"
//}
