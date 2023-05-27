package main

import (
	"user-go/db/migrations"
	"user-go/routes"
)

func main() {

	migrations.Migrations()
	routes.SetupRoutes()

}
