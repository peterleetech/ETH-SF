package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"veric-backend-mvp/logic/config"
	"veric-backend-mvp/logic/log"
)

var db *gorm.DB

func buildDSN() string {
	dbConfig := config.Get().DB
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.TimeZone,
	)
}

func mustSuccess(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDB() {
	mainDB, err := gorm.Open(mysql.Open(buildDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = mainDB
	mainDB.Logger = zapgorm2.New(log.GetLogger())

	if config.Get().Debug.Verbose {
		db = db.Debug()
	}

	migrate := db

	mustSuccess(migrate.AutoMigrate(&User{}))
	mustSuccess(migrate.AutoMigrate(&VerifiableCredential{}))
	mustSuccess(migrate.AutoMigrate(&Event{}))
	mustSuccess(migrate.AutoMigrate(&Tx{}))
	mustSuccess(migrate.AutoMigrate(&ProcessItem{}))
}
