package db

import (
	"fmt"
	"log"
	"os"

	"github.com/davidsonq/user-go/internal/configs"
	"github.com/davidsonq/user-go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectionDB() *gorm.DB {

	if os.Getenv("GOTEST") == "test" {
		db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatal(err)
		}
		db.AutoMigrate(&models.User{})

		return db
	}
	conf := configs.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s TimeZone=UTC",
		conf.DBconfigs.Host, conf.DBconfigs.User, conf.DBconfigs.Name, conf.DBconfigs.Pass, conf.DBconfigs.Port, conf.DBconfigs.Ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
