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
	router.HandleFunc("/readTweets", mdw.CheckBD(mdw.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", mdw.CheckBD(mdw.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", mdw.CheckBD(mdw.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/loadAvatar", mdw.CheckBD(routers.LoadAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", mdw.CheckBD(mdw.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/loadBanner", mdw.CheckBD(routers.LoadBanner)).Methods("GET")

	router.HandleFunc("/insertRelation", mdw.CheckBD(mdw.ValidateJWT(routers.Relation))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
