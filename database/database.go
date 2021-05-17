package database

import (
	_ "github/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)
