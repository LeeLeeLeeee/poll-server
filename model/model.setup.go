package model

import (
	"log"
	"os"

	"github.com/leeleeleeee/web-app/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global Db Connection
var Gdb *gorm.DB

type Postgres struct {
	db *gorm.DB
}

func (pg *Postgres) ConnectDatabse() {
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

	pg.db = db
}

func (pg *Postgres) MigrateDatabase() {
	db := pg.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		migrator := tx.Migrator()
		tableList := make([]interface{}, 0)
		tableList = append(tableList, &User{}, &Project{}, &Task{}, &ExampleForm{}, &ExampleLogic{}, &FormAttr{})
		tableList = append(tableList, &LogicConnection{}, &Logic{}, &Notice{}, &QuestionContent{}, &QuestionLayout{}, &QuestionLogic{})
		tableList = append(tableList, &QuestionTemplate{}, &QuestionType{}, &Question{}, &Static{})
		err := migrator.AutoMigrate(tableList...)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
		os.Exit(500)
	}
}

func (pg *Postgres) GetDB() *gorm.DB {
	return pg.db
}
