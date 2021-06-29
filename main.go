package main

import (
	"github.com/leeleeleeee/web-app/model"
)

func main() {
	db := model.ConnectDatabse()

	model.MigrateDatabase(db)
}
