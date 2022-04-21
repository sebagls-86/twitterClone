package mdw

import (
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/routers"
)

func ValidateJWT(c gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, _, _, err := routers.ProcessToken(c.GetHeader("Authorization"))
		if err != nil {
			c.Error(err)
			return
		}
		c.Next()
	}
}
