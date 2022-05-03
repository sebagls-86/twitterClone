package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sebagls-86/twitterClone/bd"
	"github.com/sebagls-86/twitterClone/jwt"
	"github.com/sebagls-86/twitterClone/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password incorrect"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) == 0 {
		http.Error(w, "Password is required", 400)
		return
	}

	document, exist := bd.LoginAttempt(t.Email, t.Password)

	if !exist {
		http.Error(w, "User or password incorrect", 400)
		return
	}

	jwtKey, err := jwt.JWTGenerator(document)

	if err != nil {
		http.Error(w, "An error ocurred trying to generate token", 400)
		return
	}

	resp := models.LoginAnswer{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
