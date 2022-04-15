package routers

import (
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func EraseRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "We need the id parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Something went wrong erasing the relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Couldn't delete relation "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	http.Error(w, "Relation borrada", 200)

}
