package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"github.com/sebagls-86/twitterClone/mdw"
	"github.com/sebagls-86/twitterClone/routers"
)

func Handlers() {

	router := gin.Default()

	router.POST("/register", mdw.CheckBD(routers.Register))
	router.POST("/login", routers.Login)

	router.GET("/profile", mdw.ValidateJWT(routers.Profile))
	router.PUT("/modifyProfile", routers.ModifyProfile)

	router.POST("/tweet", mdw.CheckBD(routers.SaveTweet))
	router.GET("/readTweets", routers.ReadTweets)
	router.DELETE("/deleteTweet", mdw.CheckBD(routers.DeleteTweet))

	// router.HandleFunc("/uploadAvatar", mdw.CheckBD(mdw.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	// router.HandleFunc("/loadAvatar", mdw.CheckBD(routers.LoadAvatar)).Methods("GET")
	// router.HandleFunc("/uploadBanner", mdw.CheckBD(mdw.ValidateJWT(routers.UploadBanner))).Methods("POST")
	// router.HandleFunc("/loadBanner", mdw.CheckBD(routers.LoadBanner)).Methods("GET")

	// router.HandleFunc("/insertRelation", mdw.CheckBD(mdw.ValidateJWT(routers.Relation))).Methods("POST")
	// router.HandleFunc("/deleteRelation", mdw.CheckBD(mdw.ValidateJWT(routers.EraseRelation))).Methods("DELETE")
	router.GET("/consultRelation", mdw.CheckBD(mdw.ValidateJWT(routers.ConsultRelation)))

	router.GET("/usersList", mdw.CheckBD(mdw.ValidateJWT(routers.UsersList)))
	// router.HandleFunc("/readTweetsFollows", mdw.CheckBD(mdw.ValidateJWT(routers.ReadTweetsRelations))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
