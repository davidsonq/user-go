package migrations

import (
	"flag"
	"fmt"
	"log"
	"user-go/db"
	"user-go/models"
)

func Migrations() {
	runMigrations := flag.Bool("migrations", false, "Executar migrações")
	flag.Parse()

	if *runMigrations {
		db := db.ConnectionDB()

		err := db.AutoMigrate(&models.User{})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Migrations run!")
	}

}
