package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/sebagls-86/twitterClone/bd"
)

func LoadAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "We need the id parameter", http.StatusBadRequest)
		return
	}

	profile, err := bd.ProfileFinder(ID)
	if err != nil {
		http.Error(w, "No user found "+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "No image found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error copying the image", http.StatusBadRequest)
		return
	}

}
