package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "Something went wrong trying to read the image ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var avatarFile string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(avatarFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading the image ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error while copying the image ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension

	status, err = bd.ChangeProfile(user, IDUser)
	if err != nil || !status {
		http.Error(w, "Error while cuploading the image to the BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
