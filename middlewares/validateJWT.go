package middlewares

import (
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/routes"
)

// ValidateJWT validates the JWT token that comes in the request
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.TokenProcess(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Invalid token" + err.Error(), http.StatusBadRequest)
			return 
		}

		next.ServeHTTP(w, r)
	}
}