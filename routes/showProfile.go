package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// ShowProfile is the route to show the profile user
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Invalid object ID", http.StatusBadRequest)
		return 
	}
	
	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "Unexpected error trying search the user" + err.Error(), http.StatusBadRequest)
		return 
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}