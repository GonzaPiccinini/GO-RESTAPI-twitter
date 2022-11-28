package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// GetFollowersTweet controller returns all tweets of the followers
func GetFollowersTweet(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page param is required", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page param must be integer and greater than 0", http.StatusBadRequest)
		return
	}

	result, status := db.GetFollowersTweets(UserID, page)
	if !status {
		http.Error(w, "Error trying read the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}