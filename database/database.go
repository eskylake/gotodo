package database

import (
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConnection *gorm.DB
)
