package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID param is required", http. StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID

	var result models.CheckRelation
	status, err := db.CheckRelation(relation)

	if err != nil || !status {
		result.Status = false
	} else {
		result.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}