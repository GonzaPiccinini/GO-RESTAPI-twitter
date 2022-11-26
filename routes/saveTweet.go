package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// SaveTweet is the controller to save a new tweet
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	
	if len(message.Message) <= 0 {
		http.Error(w, "Message si required", http.StatusBadRequest)
		return
	}

	register := models.SaveTweet{
		UserID: UserID,
		Message: message.Message,
		Date: time.Now(),
	}

	_, status, err := db.SaveTweet(register)
	if err != nil || !status {
		http.Error(w, "Unexpected error trying save the tweet" + err.Error(), http.StatusBadRequest)
		return 
	}

	w.WriteHeader(http.StatusCreated)
}