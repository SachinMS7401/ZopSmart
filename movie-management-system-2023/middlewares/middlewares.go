package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type users struct {
	username string
	password string
}

func Middleware(inner http.Handler) http.Handler {
	var user = users{
		username: "Sachin",
		password: "990299",
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials := r.Header.Get("Authorization")
		userDetails := strings.Split(credentials, " ")
		usr := strings.Split(userDetails[1], ":")

		userName, _ := base64.StdEncoding.DecodeString(usr[0])
		password, _ := base64.StdEncoding.DecodeString(usr[1])

		if string(userName) != user.username {
			err := json.NewEncoder(w).Encode("invalid username")

			if err != nil {
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if string(password) != user.password {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode("invalid password")

			if err != nil {
				return
			}
			return
		}
		inner.ServeHTTP(w, r)
	})
}
