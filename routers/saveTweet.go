package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/models"
)

func SaveTweet(w http.ResponseWriter, r *http.Request) {

	var message models.Tweet

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Something went wrong: "+err.Error(), 400)
		return
	}

	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(register)

	if err != nil {
		http.Error(w, "An error ocurred while trying to insert the register "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Something went wrong while inserting the tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	http.Error(w, "Tweet posted", 200)

}
