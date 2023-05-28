package migrations

import (
	"flag"
	"fmt"
	"log"

	"github.com/davidsonq/user-go/db"
	"github.com/davidsonq/user-go/models"
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
