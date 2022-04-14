package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sebagls-86/twitterClone/mdw"
	"github.com/sebagls-86/twitterClone/routers"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", mdw.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/login", mdw.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", mdw.CheckBD(mdw.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modifyProfile", mdw.CheckBD(mdw.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", mdw.CheckBD(mdw.ValidateJWT(routers.SaveTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
