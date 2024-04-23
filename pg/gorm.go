package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormInstance struct {
	*gorm.DB
}

func (g *gormInstance) clone() *gormInstance {
	return &gormInstance{DB: g.DB}
}

var _gormInstance *gormInstance

func NewGorm(c connectionConfig) (*gormInstance, error) {

	if _gormInstance != nil {
		return _gormInstance.clone(), nil
	}

	db, err := gorm.Open(postgres.Open(c.build()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	_gormInstance = &gormInstance{
		DB: db,
	}

	return _gormInstance.clone(), nil
}
