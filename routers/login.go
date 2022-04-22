package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/jwt"
	"github.com/sebagls-86/twitterClone/models"
)

func Login(ctx *gin.Context) {

	var t models.User

	err := ctx.ShouldBindJSON(&t)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if len(t.Email) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Email should be bigger than 0": err.Error()})
		return
	}

	document, exist := bd.LoginAttempt(t.Email, t.Password)

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"The user already exist": err.Error()})
		return
	}

	jwtKey, err := jwt.JWTGenerator(document)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Token error": err.Error()})
		return
	}

	resp := models.LoginAnswer{
		Token: jwtKey,
	}

	ctx.JSON(http.StatusOK, resp)

	ctx.SetCookie(
		"token",
		jwtKey,
		60*60*24,
		"/",
		"localhost",
		false,
		true,
	)

}
