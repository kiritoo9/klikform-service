package middlewarecomponents

import "net/http"

func Auth() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// check header authorization

			// check valid token

			// allow this token
			next(w, r)
		}
	}
}
