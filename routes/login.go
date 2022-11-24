package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/jwt"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// Login performs the login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var userModel models.User

	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		http.Error(w, "User and/or password are invalids" + err.Error(), http.StatusBadRequest)
		return 
	}
	if len(userModel.Email) == 0 {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return 
	}
	
	user, userExist := db.TryLogin(userModel.Email, userModel.Password)
	if !userExist {
		http.Error(w, "User and/or password are invalids", http.StatusBadRequest)
		return 
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Unexpected error trying generate the token", http.StatusBadRequest)
		return
	}

	response := models.LoginResponse {
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response) 

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})
}