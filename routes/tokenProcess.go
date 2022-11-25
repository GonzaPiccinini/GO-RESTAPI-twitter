package routes

import (
	"errors"
	"strings"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the user email value
var Email string

// UserID returns the models that will use in the Endpoints
var UserID string

// TokenProcess process the token and validates it
func TokenProcess(tokenStr string) (*models.Claim, bool, string, error) {
	secretKey := []byte("CBA2580MSJZ1570SM08151718")
	claims := &models.Claim{}

	splitToken := strings.Split(tokenStr, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid format token")
	}
	
	tokenStr = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err == nil {
		_, userExist, _ := db.UserExists(claims.Email)
		if userExist {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, userExist, UserID, nil
	}
	if !token.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	
	return claims, false, string(""), err
}