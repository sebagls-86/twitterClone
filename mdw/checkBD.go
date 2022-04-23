package mdw

import (
	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

var err error

func CheckBD(ctx gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		if bd.CheckConnection() == 0 {

			c.Error(err)
			return
		}

		ctx(c)
		c.Next()

	}

}
