package bootstrap

import (
	"clonecoding/internal/adapter/database"
	"clonecoding/internal/adapter/ginengine"
	"clonecoding/internal/adapter/hashing"
	"clonecoding/internal/adapter/jwt"
	"clonecoding/internal/config"
	"clonecoding/internal/domain"
	"clonecoding/internal/usecase"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
}

func InitApp() *App {
	database.InitDatabase()
	db := database.GetDatabase()
	database.InitScheme(domain.User{})
	database.InitScheme(domain.RefreshToken{})
	database.InitScheme(domain.Board{})
	database.InitScheme(domain.Post{})

	hashing := &hashing.HashingImpl{}

	jwtService := &jwt.JWTServiceImpl{SecretKey: []byte(config.JWTSecret)}

	userRepo := &database.UserRepositoryImpl{DB: db}
	authRepo := &database.AuthRepositoryImpl{DB: db}
	boardRepo := &database.BoardRepositoryImpl{DB: db}
	postRepo := &database.PostRepositoryImpl{DB: db}

	userUsecase := &usecase.UserUsecase{UserRepo: userRepo, Hashing: hashing}
	authUsecase := &usecase.AuthUseCase{UserRepo: userRepo, AuthRepo: authRepo, JWTService: jwtService, Hashing: hashing}
	boardUsecase := &usecase.BoardUsecase{BoardRepo: boardRepo}
	postUsecase := &usecase.PostUsecase{PostRepo: postRepo}

	userHandler := &ginengine.UserHandler{UserUseCase: userUsecase}
	authHandler := &ginengine.AuthHandler{AuthUseCase: authUsecase}
	boardHandler := &ginengine.BoardHandler{BoardUseCase: boardUsecase}
	postHandler := &ginengine.PostHandler{PostUsecase: postUsecase}

	routerDeps := ginengine.RouterDeps{
		UserHandler:  userHandler,
		AuthHandler:  authHandler,
		BoardHandler: boardHandler,
		PostHandler:  postHandler,
		JwtService:   jwtService,
	}
	r := ginengine.SetRouter(&routerDeps)

	app := App{Engine: r}

	return &app
}

func (a *App) Run() {
	a.Engine.Run(":" + config.Port)
}
