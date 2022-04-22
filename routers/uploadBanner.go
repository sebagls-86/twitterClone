package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func UploadBanner(ctx *gin.Context) {

	file, handler, err := ctx.Request.FormFile("banner")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Something went wrong trying to read the image ": err.Error()})
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var avatarFile string = "uploads/banners/" + IDUser + "." + extension

	f, err := os.OpenFile(avatarFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error while uploading the image ! ": err.Error()})
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error while copying the image ! ": err.Error()})
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension

	status, err = bd.ChangeProfile(user, IDUser)
	if err != nil || !status {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error while cuploading the image to the BD ! ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Upload successfully"})

}
