package routes

import (
	"fmt"

	"github.com/davidsonq/user-go/configs"
	"github.com/davidsonq/user-go/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	userRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(fmt.Sprintf(":%s", configs.GetConfig().APIconfigs.Port))

	return r
}

func userRoutes(r *gin.Engine) {
	r.POST("/api/users", handlers.CreateUserHandle)
}
