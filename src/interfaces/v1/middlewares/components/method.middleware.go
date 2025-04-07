package middlewarecomponents

import (
	"encoding/json"
	"net/http"
)

func Method(method []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			allowed := false

			// check if method is allowed
			for _, item := range method {
				if item == r.Method {
					allowed = true
					break
				}
			}

			// handle response
			if allowed == false {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"message": "Method is not allowed",
				})
				return
			}
			next(w, r)
		}
	}
}
