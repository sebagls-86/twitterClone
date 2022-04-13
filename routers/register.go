package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Mail is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password minimum lenght 6 characters", 400)
		return
	}

	_, userFound, _ := bd.CheckUserExist(t.Email)

	if userFound {
		http.Error(w, "Email already exist", 400)
		return
	}

	_, status, err := bd.NewRegister(t)

	if err != nil {
		http.Error(w, "Something went wrong trying to registe the user: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Something went wrong while registering the user: "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
