package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonyangmh/sayakaya/dto"
	"github.com/jasonyangmh/sayakaya/shared"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *shared.CustomError:
				c.AbortWithStatusJSON(e.Code, dto.JSONResponse{Error: e})
				return
			default:
				c.AbortWithStatusJSON(shared.ErrUnknownError.Code, dto.JSONResponse{Error: shared.ErrUnknownError.Error()})
			}
		}
	}
}
