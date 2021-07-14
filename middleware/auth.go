package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthCookie(c *gin.Context) {
	defer func() {
		token, err := c.Cookie("token")
		if err != nil {
			print(err.Error())
		}
		print(token)
	}()
	c.Next()
}
