package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func ConsultRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "We need the id parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.ResponseConsultRelation

	status, err := bd.CheckRelation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
