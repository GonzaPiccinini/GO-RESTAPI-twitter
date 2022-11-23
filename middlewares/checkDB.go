package middlewares

import (
	"net/http"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/db"
)

// CheckDB checks if the database connection is up
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "CONEXION PERDIDA CON LA BASE DE DATOS", 500)
			return 
		}
		next.ServeHTTP(w, r)
	}
}