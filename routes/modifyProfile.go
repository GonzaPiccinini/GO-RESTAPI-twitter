package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// ModifyProfile is the controller that allows modify a user profile
func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid data" + err.Error(), http.StatusBadRequest)
		return 
	}

	var status bool
	status, err = db.ModifyUser(user, UserID)
	if err != nil || !status {
		http.Error(w, "Unexpected error trying modify profile" + err.Error(), http.StatusBadRequest)
		return 
	} 

	w.WriteHeader(http.StatusCreated)
}