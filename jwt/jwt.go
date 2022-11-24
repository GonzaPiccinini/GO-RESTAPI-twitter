package jwt

import (
	"time"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a new JSON Web Token
func GenerateJWT(userModel models.User) (string, error) {
	secretKey := []byte("CBA2580MSJZ1570SM08151718")
	payload := jwt.MapClaims{
		"email":     userModel.Email,
		"name":      userModel.Name,
		"lastname":  userModel.Lastname,
		"birthday":  userModel.Birthday,
		"biography": userModel.Biography,
		"location":  userModel.Location,
		"website":   userModel.Website,
		"_id":       userModel.ID.Hex(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return tokenStr, err
	}
	
	return tokenStr, nil
}