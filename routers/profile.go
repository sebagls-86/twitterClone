package routers

import (
	//"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func Profile(ctx *gin.Context) {

	var err error

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the ID": err.Error()})
		return
	}

	profile, err := bd.ProfileFinder(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"We cant find the user": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profile)

}
