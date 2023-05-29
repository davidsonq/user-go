package routes

import (
	"fmt"

	"github.com/davidsonq/user-go/internal/configs"
	"github.com/davidsonq/user-go/internal/handlers"
	"github.com/davidsonq/user-go/internal/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	userRoutes(r)
	LoginRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(fmt.Sprintf(":%s", configs.GetConfig().APIconfigs.Port))

	return r
}

func userRoutes(r *gin.Engine) {
	r.POST("/api/users", handlers.CreateUserHandle)
	r.GET("/api/users/profile", middlewares.AuthMiddleware(), handlers.GetProfileUser)
	// r.PATCH("/api/users/{id}", middlewares.AuthMiddleware(), handlers.GetProfileUser)
	// r.DELETE("/api/users/{id}", middlewares.AuthMiddleware(), handlers.GetProfileUser)
}

func LoginRoutes(r *gin.Engine) {
	r.POST("/api/users/login", handlers.LoginUserHandler)
}
