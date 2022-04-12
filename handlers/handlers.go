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

	router.HandleFunc("/registro", mdw.CheckBD(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
