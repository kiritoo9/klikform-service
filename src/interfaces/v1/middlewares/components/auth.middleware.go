package middlewarecomponents

import (
	"context"
	"encoding/json"
	"fmt"
	repos "klikform/src/applications/repos/masters"
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
			token, err := jwt.ParseWithClaims(authorization, claims, func(token *jwt.Token) (any, error) {
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

			// send user logged data into request context
			var role string
			userID := (*claims)["id"].(string)
			userRole, err := repos.GetRoleByUser(userID)
			if err == nil {
				role = userRole.Role.Name
			}

			userClaims := map[string]any{
				"id":   userID,
				"role": role,
			}
			ctx := context.WithValue(r.Context(), "loggedToken", userClaims)
			next(w, r.WithContext(ctx)) // pass with context
		}
	}
}
