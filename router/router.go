package router

import "github.com/gin-gonic/gin"

func SetRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("")
		}
	}
}
