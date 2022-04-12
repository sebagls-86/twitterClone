package mdw

import (
	"net/http"

	"github.com/sebagls-86/twitterClone/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
