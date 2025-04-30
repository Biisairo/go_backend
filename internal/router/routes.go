package router

import (
	"clonecoding/internal/adapter/http"
	"clonecoding/internal/adapter/jwt"

	"github.com/gin-gonic/gin"
)

func SetRouter(
	userHandler *http.UserHandler,
	authHandler *http.AuthHandler,
	jwtService *jwt.JWTServiceImpl,
) *gin.Engine {
	r := gin.Default()

	r.POST("/create", userHandler.CreateUser)

	user := r.Group("/user")
	user.Use(http.JWTMiddleware(jwtService))
	user.GET("/", userHandler.GetAllUser)
	user.GET("/:id", userHandler.GetUser)

	auth := r.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh", authHandler.Refresh)
	auth.POST("/logout", authHandler.Logout)

	return r
}
