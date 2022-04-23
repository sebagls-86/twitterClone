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

	router.POST("/register", routers.Register)
	router.POST("/login", routers.Login)

	router.GET("/profile", mdw.ValidateJWT(routers.Profile))

	router.PUT("/modifyProfile", mdw.CheckBD(mdw.ValidateJWT(routers.ModifyProfile)))
	router.POST("/tweet", mdw.CheckBD(mdw.ValidateJWT(routers.SaveTweet)))
	router.GET("/readTweets", mdw.CheckBD(mdw.ValidateJWT(routers.ReadTweets)))
	router.DELETE("/deleteTweet", mdw.CheckBD(mdw.ValidateJWT(routers.DeleteTweet)))

	router.POST("/uploadAvatar", mdw.CheckBD(mdw.ValidateJWT(routers.UploadAvatar)))
	router.GET("/loadAvatar", routers.LoadAvatar)
	router.POST("/uploadBanner", mdw.CheckBD(mdw.ValidateJWT(routers.UploadBanner)))
	router.GET("/loadBanner", routers.LoadBanner)

	router.POST("/insertRelation", mdw.CheckBD(mdw.ValidateJWT(routers.Relation)))
	router.DELETE("/deleteRelation", mdw.CheckBD(mdw.ValidateJWT(routers.EraseRelation)))
	router.GET("/consultRelation", mdw.CheckBD(mdw.ValidateJWT(routers.ConsultRelation)))

	router.GET("/usersList", mdw.CheckBD(mdw.ValidateJWT(routers.UsersList)))
	router.GET("/readTweetsFollows", mdw.CheckBD(mdw.ValidateJWT(routers.ReadTweetsRelations)))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
