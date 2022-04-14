package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sebagls-86/twitterClone/bd"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "We need the id parameter", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "We need the page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "The page parameter must be greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	response, correct := bd.ReadTweets(ID, pag)
	if !correct {
		http.Error(w, "Error reading twwets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
