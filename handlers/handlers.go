package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/middlewares"
	"github.com/GonzaPiccinini/GO-RESTAPI-twitter/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers create the router object, set the PORT and set CORS
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routes.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}