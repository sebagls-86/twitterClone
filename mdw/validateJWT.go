package mdw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/routers"
)

func ValidateJWT(ctx gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		_, _, _, err := routers.ProcessToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx(c)
		c.Next()

	}

}
