package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

var err error

func ReadTweets(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the id parameter": err.Error()})
		return
	}

	if len(ctx.Query("page")) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the page parameter": err.Error()})
		return
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"The page parameter must be greater than 0": err.Error()})
		return
	}

	pag := int64(page)

	response, correct := bd.ReadTweets(ID, pag)
	if !correct {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error reading tweets": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
