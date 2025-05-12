package ginengine

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, APIResponse{
		Data: data,
	})
}
func Fail(c *gin.Context, status int, err string) {
	c.JSON(status, APIResponse{
		Error: err,
	})
}
