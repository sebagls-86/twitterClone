package mdw

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/routers"
)

func ValidateJWT(c gin.HandlerFunc) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		_, _, _, err := routers.ProcessToken(ctx.GetHeader("Authorization"))
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusOK, gin.H{"message": "paso por ValidateJWT con error"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "paso por ValidateJWT"})
		ctx.Request.Context()

	}

}
