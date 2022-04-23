package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func DeleteTweet(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "We need the id parameter"})
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"An error ocurred while deleting the tweet": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})

}
