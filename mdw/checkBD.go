package mdw

import (
	"github.com/gin-gonic/gin"
	"github.com/sebagls-86/twitterClone/bd"
)

var err error

func CheckBD(c gin.HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {

		if bd.CheckConnection() == 0 {

			c.Error(err)
			//panic(c.Error(err))
			return
		}

	}

}

// func CheckBD(next http.HandlerFunc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if bd.CheckConnection() == 0 {
// 			http.Error(w, "Connection lost", 500)
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	}
// }
