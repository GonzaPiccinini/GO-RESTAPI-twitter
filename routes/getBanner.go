package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// GetBanner controller gets banner from user
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 { 
		http.Error(w, "ID param si required", http.StatusBadRequest)
		return 
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error trying copy the image", http.StatusBadRequest)
	}
}