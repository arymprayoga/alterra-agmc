package migration

import (
	"day6/internal/model"

	"day6/database"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
