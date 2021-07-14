package middleware

import (
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"tiga/model"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err := convertError(r)
			c.JSON(err.Code, err.Msg)
			c.Abort()
		}
	}()
	c.Next()
}

func convertError(r interface{}) model.HttpError {
	switch v := r.(type) {
	case model.HttpError:
		return v
	case error:
		err := model.UnknownServerError
		err.Err = v
		return err
	default:
		return model.UnknownServerError
	}
}
