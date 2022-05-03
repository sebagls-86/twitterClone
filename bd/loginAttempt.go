package bd

import (
	"github.com/sebagls-86/twitterClone/models"
	"golang.org/x/crypto/bcrypt"
)

// It takes an email and password, checks if the user exists, and if it does, it compares the password
// with the one in the database
func LoginAttempt(email string, password string) (models.User, bool) {

	us, found, _ := CheckUserExist(email)
	if !found {
		return us, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(us.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return us, false
	}

	return us, true
}
