package tests

import (
	"clonecoding/config"
	"clonecoding/database"
	userDatabase "clonecoding/internal/adapter/database"
	"clonecoding/internal/adapter/hashing"
	"clonecoding/internal/adapter/http"
	"clonecoding/internal/adapter/jwt"
	"clonecoding/internal/domain"
	"clonecoding/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupTestEnv() *gin.Engine {
	gin.SetMode(gin.TestMode)

	os.Chdir("..")

	config.LoadConfig()

	database.InitDatabase()
	db := database.GetDatabase()
	database.InitScheme(domain.User{})
	database.InitScheme(domain.RefreshToken{})

	hashing := &hashing.HashingImpl{}

	jwtService := &jwt.JWTServiceImpl{SecretKey: []byte(config.JWTSecret)}

	userRepo := &userDatabase.UserRepositoryImpl{DB: db}
	authRepo := &userDatabase.AuthRepositoryImpl{DB: db}

	userUseCase := &usecase.UserUsecase{UserRepo: userRepo, Hashing: hashing}
	userHandler := &http.UserHandler{UserUseCase: userUseCase}

	authUsecase := &usecase.AuthUseCase{UserRepo: userRepo, AuthRepo: authRepo, JWTService: jwtService, Hashing: hashing}
	authHandler := &http.AuthHandler{AuthUseCase: authUsecase}

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
