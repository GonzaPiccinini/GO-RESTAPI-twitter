package routes

import (
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// DeleteTweet controller deletes a specific tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id param is required", http.StatusBadRequest)
		return 
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "Unexpected error trying delete the tweet" + err.Error(), http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}