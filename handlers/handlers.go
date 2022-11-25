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
	router.HandleFunc("/login", middlewares.CheckDB(routes.Login)).Methods("POST")
	router.HandleFunc("/showprofile", middlewares.CheckDB(middlewares.ValidateJWT(routes.ShowProfile))).Methods("GET")
	router.HandleFunc("/modifyprofile", middlewares.CheckDB(middlewares.ValidateJWT(routes.ModifyProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}