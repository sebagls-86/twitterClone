package mdw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/routers"
)

func ValidateJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		_, _, _, err := routers.ProcessToken(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"We cant find the user": err.Error()})
			//panic("We cant find the user")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "paso por JWT"})
		c.Next()

	}
}
