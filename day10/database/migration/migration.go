package migration

import (
	"day10/internal/model"

	"day10/database"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
