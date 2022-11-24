package db

import (
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin returns the user if the login attempt is ok
func TryLogin(email string, password string) (models.User, bool) {
	user, userExist, _ := UserExists(email)

	if !userExist {
		return user, false
	}

	passwordRequest := []byte(password)
	encryptedPasswordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(encryptedPasswordDB, passwordRequest)

	if err != nil {
		return user, false
	}

	return user, true
}