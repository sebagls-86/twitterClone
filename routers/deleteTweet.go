package routers

import (
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "We need the id parameter", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "An error ocurred while deleting the tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	http.Error(w, "Tweet borrado con exito", 200)

}
