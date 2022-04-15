package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sebagls-86/twitterClone/bd"
)

func ReadTweetsRelations(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "We need the page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "We need the page parameter in numbers", http.StatusBadRequest)
		return
	}

	response, correct := bd.ReadTweetsFollows(IDUser, page)
	if !correct {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
