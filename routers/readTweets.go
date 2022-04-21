package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func ReadTweets(ctx *gin.Context) {

	var err error

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the user ID": err.Error()})
		return
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Page missing": err.Error()})
		return
	}

	pag := int64(page)

	response, correct := bd.ReadTweets(ID, pag)
	if !correct {
		ctx.JSON(http.StatusBadRequest, gin.H{"Soemthing went wrong": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
