package migration

import (
	"day7/internal/model"

	"day7/database"
)

var tables = []interface{}{
	&model.User{},
}

func Migrate() {
	conn := database.GetConnection()
	conn.AutoMigrate(tables...)
}
