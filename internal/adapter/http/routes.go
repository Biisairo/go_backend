package http

import (
	"clonecoding/internal/adapter/jwt"

	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	UserHandler  *UserHandler
	AuthHandler  *AuthHandler
	BoardHandler *BoardHandler
	PostHandler  *PostHandler
	JwtService   *jwt.JWTServiceImpl
}

func SetRouter(routerDeps *RouterDeps) *gin.Engine {
	r := gin.Default()

	r.POST("/create", routerDeps.UserHandler.CreateUser)

	auth := r.Group("/auth")
	{
		auth.POST("/login", routerDeps.AuthHandler.Login)
		auth.POST("/refresh", routerDeps.AuthHandler.Refresh)
		auth.POST("/logout", routerDeps.AuthHandler.Logout)
	}

	user := r.Group("/user")
	userAuth := user.Group("")
	userAuth.Use(JWTMiddleware(routerDeps.JwtService))
	{
		userAuth.GET("/", routerDeps.UserHandler.GetAllUser)
		userAuth.GET("/:user_id", routerDeps.UserHandler.GetUser)
		userAuth.GET("/post/:user_id", routerDeps.PostHandler.GetPostByUserId) // get post data of specific user
	}

	board := r.Group("/board")
	boardAuth := board.Group("")
	boardAuth.Use(JWTMiddleware(routerDeps.JwtService))
	{
		boardAuth.POST("/", routerDeps.BoardHandler.CreateBoard)
		boardAuth.GET("/", routerDeps.BoardHandler.GetAllBoard)
		boardAuth.GET("/:board_id", routerDeps.BoardHandler.GetBoard)             // get specific board data
		boardAuth.POST("/post/:board_id", routerDeps.PostHandler.CreatePost)      // get post data of specific board
		boardAuth.GET("/post/:board_id", routerDeps.PostHandler.GetPostByBoardId) // get post data of specific board
	}

	post := r.Group("/post")
	postAuth := post.Group("")
	postAuth.Use(JWTMiddleware(routerDeps.JwtService))
	{
		postAuth.GET("/", routerDeps.PostHandler.GetAllPost) // get all post data
	}

	return r
}
