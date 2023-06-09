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
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})

	userRoutes(r)
	LoginRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(fmt.Sprintf(":%s", configs.GetConfig().APIconfigs.Port))

	return r
}

func userRoutes(r *gin.Engine) {
	r.POST("/api/users", handlers.CreateUserHandle)
	r.GET("/api/users/profile", middlewares.AuthMiddleware(), handlers.GetProfileUserHandler)
	r.PATCH("/api/users/:id", middlewares.AuthMiddleware(), handlers.UpdatedUserHandle)
	r.DELETE("/api/users/:id", middlewares.AuthMiddleware(), handlers.DeleteUserHandler)
}

func LoginRoutes(r *gin.Engine) {
	r.POST("/api/users/login", handlers.LoginUserHandler)
}
