package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func UsersList(ctx *gin.Context) {

	typeUser := ctx.Query("type")
	page := ctx.Query("page")
	searchType := ctx.Query("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Parameter page must be greater than 0"})
		return
	}

	pag := int64(pagTemp)

	result, status := bd.ReadAllUsers(IDUser, pag, searchType, typeUser)
	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Error reading users"})
		return
	}

	ctx.JSON(http.StatusOK, result)
	ctx.JSON(http.StatusOK, gin.H{"message": "Here I am"})

}
