package routes

import (
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// HighRelation controller performs the relation between users
func HighRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID param is required", http.StatusBadRequest)
		return 
	}

	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID
	
	status, err := db.CreateRelation(relation)
	if err != nil || !status {
		http.Error(w, "Unexpected error trying make relation" + err.Error(), http.StatusBadRequest)
		return 
	}

	w.WriteHeader(http.StatusCreated)
}