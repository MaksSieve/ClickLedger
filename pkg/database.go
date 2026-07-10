package db

import "gorm.io/gorm"

type Database interface {
	MakeConnection() (*gorm.DB, error)
}
