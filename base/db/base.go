package db

import (
	"gorm.io/gorm"
)

type DBSelector interface {
	Create() (*gorm.DB, error)
}

type DBBase struct {
	DbType DBSelector
}


