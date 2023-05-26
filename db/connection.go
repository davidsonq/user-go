package db

import (
	"fmt"
	"log"
	"os"
	"user-go/configs"
	"user-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectionDB() *gorm.DB {
	conf := configs.GetConfig()

	if os.Getenv("GOTEST") == "test" {
		db, err := gorm.Open(sqlite.Open("flag::memory:?cache=false"), &gorm.Config{
			Logger: logger.Discard.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatal(err)
		}
		db.AutoMigrate(&models.User{})

		return db
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=require TimeZone=UTC",
		conf.DBconfigs.Host, conf.DBconfigs.User, conf.DBconfigs.Name, conf.DBconfigs.Pass, conf.DBconfigs.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
