package middleware

import (
	"github.com/gin-gonic/gin"
	"tiga/model"
	"tiga/util"
)

func AuthCookie(c *gin.Context) {
	defer func() {
		token, err := c.Cookie("token")
		if err != nil {
			panic(model.UnauthorizedError)
		}
		claim, ok := util.CheckToken(token)
		if !ok {
			c.Set("user", claim)
			panic(model.UnexpectedTokenError)
		}
	}()
	c.Next()
}
