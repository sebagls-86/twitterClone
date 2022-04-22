package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

func LoadAvatar(ctx *gin.Context) {

	ID := ctx.Query("id")

	if len(ID) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"We need the id parameter": err.Error()})
		return
	}

	profile, err := bd.ProfileFinder(ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No user found ": err.Error()})
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No image found ": err.Error()})
		return
	}

	_, err = io.Copy(ctx.Writer, OpenFile)
	if err != nil {
		http.Error(ctx.Writer, "Error copying the image", http.StatusBadRequest)
		return
	}

}
