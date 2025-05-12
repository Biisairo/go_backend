package ginengine

import (
	"clonecoding/internal/adapter/jwt"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(j *jwt.JWTServiceImpl) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			Fail(c, http.StatusUnauthorized, "error: Missing or invalid token")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := j.ParseToken(tokenStr)
		if err != nil {
			Fail(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		fmt.Println(claims.ExpiresAt)
		fmt.Println(time.Now())
		if claims.ExpiresAt.Before(time.Now()) {
			Fail(c, http.StatusUnauthorized, "expired token")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
