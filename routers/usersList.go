package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sebagls-86/twitterClone/bd"
)

func UsersList(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	searchType := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Parameter page must be greater than 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.ReadAllUsers(IDUser, pag, searchType, typeUser)
	if !status {
		http.Error(w, "Error reading users "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
