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
	auth.POST("/login", routerDeps.AuthHandler.Login)
	auth.POST("/refresh", routerDeps.AuthHandler.Refresh)
	auth.POST("/logout", routerDeps.AuthHandler.Logout)

	user := r.Group("/user")
	user.Use(JWTMiddleware(routerDeps.JwtService))
	user.GET("/", routerDeps.UserHandler.GetAllUser)
	user.GET("/:id", routerDeps.UserHandler.GetUser)
	user.GET("/post/:userId", routerDeps.PostHandler.GetPostByUserId) // get post data of specific user

	board := r.Group("/board")
	board.POST("/", routerDeps.BoardHandler.CreateBoard)
	board.GET("/", routerDeps.BoardHandler.GetAllBoard)
	board.GET("/:id", routerDeps.BoardHandler.GetBoard)                  // get specific board data
	board.POST("/post/:boardId", routerDeps.PostHandler.CreatePost)      // get post data of specific board
	board.GET("/post/:boardId", routerDeps.PostHandler.GetPostByBoardId) // get post data of specific board

	post := r.Group("/post")
	post.GET("/", routerDeps.PostHandler.GetAllPost) // get all post data

	return r
}
