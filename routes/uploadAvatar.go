package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
)

// UploadAvatar controller allos save the user avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	fileForm, handler, _ := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var file string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(file, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error trying up the file" + err.Error(), http.StatusBadRequest)
		return 
	}

	_, err = io.Copy(f, fileForm)
	if err != nil {
		http.Error(w, "Error trying up the file" + err.Error(), http.StatusBadRequest)
		return 
	}

	var user models.User
	var status bool
	
	user.Avatar = UserID + "." + extension
	status, err = db.ModifyUser(user, UserID)
	if err != nil || !status {
		http.Error(w, "Error trying up the file" + err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}