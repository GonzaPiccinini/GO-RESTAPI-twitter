package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// GetTweets controller allows returns a user's tweets
func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id param is required", http.StatusBadRequest)
		return 
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page param is required", http.StatusBadRequest)
		return 
	}
	pageStr, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if pageStr < 1 {
		http.Error(w, "page param must be greater than 0", http.StatusBadRequest)
		return 
	}
	page := int64(pageStr)

	result, status := db.GetTweets(ID, page)
	if !status {
		http.Error(w, "Unexpected error trying get tweets", http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}