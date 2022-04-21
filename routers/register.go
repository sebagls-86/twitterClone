package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func Register(ctx *gin.Context) {
	var t models.User
	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		fmt.Println("paso por aca 1")
		return
	}

	if len(t.Email) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}
	if len(t.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	_, userFound, _ := bd.CheckUserExist(t.Email)

	if userFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	_, status, err := bd.NewRegister(t)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	if !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})

}
