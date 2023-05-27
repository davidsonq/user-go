package routes

import (
	"fmt"
	"user-go/configs"
	"user-go/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	userRoutes(r)

	r.Run(fmt.Sprintf(":%s", configs.GetConfig().APIconfigs.Port))

}

func userRoutes(r *gin.Engine) {
	r.POST("/api/users", handlers.CreateUser)
}
