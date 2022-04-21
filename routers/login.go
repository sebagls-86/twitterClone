package routers

import (
	//"encoding/json"
	"net/http"

	//"time"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	jwt "github.com/sebagls-86/twitterClone/jwt"
	"github.com/sebagls-86/twitterClone/models"
)

func Login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	if len(user.Email) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Email is required": err.Error()})
		return
	}

	document, exist := bd.LoginAttempt(user.Email, user.Password)

	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"User or password incorrect": err.Error()})
		return
	}

	jwtKey, err := jwt.JWTGenerator(document)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"An error ocurred trying to generate token": err.Error()})
		return
	}

	//var w http.ResponseWriter

	resp := models.LoginAnswer{
		Token: jwtKey,
	}

	ctx.JSON(http.StatusOK, resp)
	//json.NewEncoder(w).Encode(resp)

	//expirationTime := time.Now().Add(24 * time.Hour)

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
