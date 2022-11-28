package routes

import (
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

func LowRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID param is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID
	
	status, err := db.DeleteRelation(relation)
	if err != nil || !status {
		http.Error(w, "Unexpected error trying delete relation", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}