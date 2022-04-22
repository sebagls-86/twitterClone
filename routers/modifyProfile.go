package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func ModifyProfile(ctx *gin.Context) {

	var t models.User

	err := ctx.ShouldBindJSON(&t)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	var status bool
	status, err = bd.ChangeProfile(t, IDUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"An error ocurred while modifying the register": err.Error()})
		return
	}

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"omething went wrong trying to modify the register": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Profile changed"})

}
