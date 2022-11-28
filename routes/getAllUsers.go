package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// GetAllUsers controller returns a list of users according to the query
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	pageQuery := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(pageQuery)
	if err != nil {
		http.Error(w, "page param is required and must be integer and greater than 0", http.StatusBadRequest)
		return
	}

	page := int64(pageTemp)

	results, status := db.GetAllUsers(UserID, page, search, typeUser)
	if !status {
		http.Error(w, "page param is required and must be integer and greater than 0", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)
}