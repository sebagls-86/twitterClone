package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func SaveTweet(ctx *gin.Context) {

	var message models.Tweet

	err := ctx.ShouldBindJSON(&message)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Something went wrong": err.Error()})
		return
	}

	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"An error ocurred while trying to insert the register": err.Error()})
		return
	}
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"Something went wrong while inserting the tweet": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tweet posted"})

}
