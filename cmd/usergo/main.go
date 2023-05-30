package main

import (
	docs "github.com/davidsonq/user-go/docs"
	"github.com/davidsonq/user-go/internal/db/migrations"
	"github.com/davidsonq/user-go/internal/routes"
)

// @title           User Management Microservice with Login System
// @version         1.0
// @description     This is a microservice built to manage users, with authentication and login features.

// @contact.email  davidsonquaresma@hotmail.com

// @license.name  MIT
// @license.url   https://github.com/davidsonq/user-go/blame/main/LICENSE

// @host      usergolang.onrender.com
// @BasePath  /api/
func main() {
	migrations.Migrations()
	routes.SetupRoutes()
	docs.SwaggerInfo.BasePath = "/"

}
