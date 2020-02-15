package middleware

import (
	"app/config"
	"app/handler/user"
	"app/initDB"
	"github.com/gin-gonic/gin"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !initDB.TokenExist(user.Token) {
			c.Abort()
			c.String(200, config.ErrorReasonUnLogin)
		}
	}
}
