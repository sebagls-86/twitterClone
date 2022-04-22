package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func ReadTweetsRelations(ctx *gin.Context) {

	if len(ctx.Query("page")) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the page parameter": err.Error()})
		return
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the page parameter in numbers": err.Error()})
		return
	}

	response, correct := bd.ReadTweetsFollows(IDUser, page)
	if !correct {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error reading tweets": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)

}
