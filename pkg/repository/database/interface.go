package database

import "gorm.io/gorm"

type Client interface {
	FindOneBy(key, value string, entity interface{})
	FindOneById(id uint, entity interface{}) interface{}
	Save(entity interface{})
	Migrate(entity interface{})
	Update(payload, entity interface{})
	MigrateAndSeed(entity interface{}, payload ...interface{})
	FindAllByPage(entity interface{}, page int, pageNumber *int64)
	GetGORMClient() *gorm.DB
}
