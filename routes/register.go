package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// Register is the function to create a user in the database
func Register(w http.ResponseWriter, r *http.Request) {
	var userModel models.User

	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		http.Error(w, "Validation failed" + err.Error(), http.StatusBadRequest)
		return
	}

	if len(userModel.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	if len(userModel.Password) < 6 {
		http.Error(w, "Password length must be 6 or than", http.StatusBadRequest)
		return
	}

	_, userExist, _ := db.UserExists(userModel.Email)
	if userExist {
		http.Error(w, "The user already exists", http.StatusBadRequest)
		return 
	}

	_, status, err := db.RegisterUser(userModel)
	if err != nil || !status {
		http.Error(w, "Unexpected error trying register user" + err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusContinue)
}