package middlewarecomponents

import (
	"encoding/json"
	"fmt"
	"klikform/src/infras/configs"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func Auth() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// check header authorization
			authorization := r.Header.Get("authorization")
			if authorization == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "Missing bearer header",
				})
				return
			}

			// check valid token
			configs := configs.LoadConfig()
			claims := &jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("done")
				}
				return configs.JWT_SECRET, nil
			})

			if err != nil || !token.Valid {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]any{
					"message": "Token authorization is not valid",
				})
				return
			}

			// allow this token
			next(w, r)
		}
	}
}
