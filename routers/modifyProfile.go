package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password incorrect"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ChangeProfile(t, IDUser)
	if err != nil {
		http.Error(w, "An error ocurred while modifying the register"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Something went wrong trying to modify the register"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
