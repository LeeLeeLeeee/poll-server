package model

import (
	"log"
	"os"

	"github.com/leeleeleeee/web-app/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabse() *gorm.DB {
	configuration := conf.GetConfig()
	connectInfo := configuration.DbConnect

	dsn := "host="
	dsn += connectInfo.Host
	dsn += " user="
	dsn += connectInfo.User
	dsn += " password="
	dsn += connectInfo.Password
	dsn += " dbname="
	dsn += connectInfo.DatabaseName
	dsn += " port="
	dsn += connectInfo.Port
	dsn += " sslmode="
	dsn += connectInfo.Sslmode
	dsn += " TimeZone="
	dsn += connectInfo.TimeZone

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		os.Exit(500)
	}
	db.Migrator()

	return db
}

func MigrateDatabase(db *gorm.DB) {
	migrator := db.Migrator()
	tableList := make([]interface{}, 0)
	tableList = append(tableList, &User{})
	migrator.CreateTable(tableList...)
}
