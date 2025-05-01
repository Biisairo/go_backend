package bootstrap

import (
	"clonecoding/internal/adapter/database"
	"clonecoding/internal/adapter/hashing"
	"clonecoding/internal/adapter/http"
	"clonecoding/internal/adapter/jwt"
	"clonecoding/internal/config"
	"clonecoding/internal/domain"
	"clonecoding/internal/router"
	"clonecoding/internal/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func InitApp() *App {
	config.LoadConfig()

	database.InitDatabase()
	db := database.GetDatabase()
	database.InitScheme(domain.User{})
	database.InitScheme(domain.RefreshToken{})
	database.InitScheme(domain.Board{})

	hashing := &hashing.HashingImpl{}

	jwtService := &jwt.JWTServiceImpl{SecretKey: []byte(config.JWTSecret)}

	userRepo := &database.UserRepositoryImpl{DB: db}
	authRepo := &database.AuthRepositoryImpl{DB: db}
	boardRepo := &database.BoardRepositoryImpl{DB: db}

	userUsecase := &usecase.UserUsecase{UserRepo: userRepo, Hashing: hashing}
	authUsecase := &usecase.AuthUseCase{UserRepo: userRepo, AuthRepo: authRepo, JWTService: jwtService, Hashing: hashing}
	boardUsecase := &usecase.BoardUsecase{BoardRepo: boardRepo}

	userHandler := &http.UserHandler{UserUseCase: userUsecase}
	authHandler := &http.AuthHandler{AuthUseCase: authUsecase}
	boardHandler := &http.BoardHandler{BoardUseCase: boardUsecase}

	r := router.SetRouter(userHandler, authHandler, boardHandler, jwtService)

	app := App{Engine: r}

	return &app
}

func (a *App) Run() {
	a.Engine.Run(":" + config.Port)
}
