package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
)

func Profile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "We need an ID number", http.StatusBadRequest)
		return
	}

	profile, err := bd.ProfileFinder(ID)
	if err != nil {
		http.Error(w, "Error trying to find the profile "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
