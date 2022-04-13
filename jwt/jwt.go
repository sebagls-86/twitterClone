package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sebagls-86/twitterClone/models"
)

func JWTGenerator(t models.User) (string, error) {
	myKey := []byte("SecretKey")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDay,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
