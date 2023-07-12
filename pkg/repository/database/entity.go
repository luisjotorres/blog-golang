package database

import "gorm.io/gorm"

type client struct {
	*gorm.DB
}
