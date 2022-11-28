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
	router.HandleFunc("/showProfile", middlewares.CheckDB(middlewares.ValidateJWT(routes.ShowProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlewares.CheckDB(middlewares.ValidateJWT(routes.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/createTweet", middlewares.CheckDB(middlewares.ValidateJWT(routes.SaveTweet))).Methods("POST")
	router.HandleFunc("/getTweets", middlewares.CheckDB(middlewares.ValidateJWT(routes.GetTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlewares.CheckDB(middlewares.ValidateJWT(routes.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routes.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routes.GetAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routes.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDB(middlewares.ValidateJWT(routes.GetBanner))).Methods("GET")
	router.HandleFunc("/highRelation", middlewares.CheckDB(middlewares.ValidateJWT(routes.HighRelation))).Methods("POST")
	router.HandleFunc("/lowRelation", middlewares.CheckDB(middlewares.ValidateJWT(routes.LowRelation))).Methods("DELETE")
	router.HandleFunc("/checkRelation", middlewares.CheckDB(middlewares.ValidateJWT(routes.CheckRelation))).Methods("GET")
	router.HandleFunc("/getAllUsers", middlewares.CheckDB(middlewares.ValidateJWT(routes.GetAllUsers))).Methods("GET")
	router.HandleFunc("/getFollowersTweets", middlewares.CheckDB(middlewares.ValidateJWT(routes.GetFollowersTweet))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}