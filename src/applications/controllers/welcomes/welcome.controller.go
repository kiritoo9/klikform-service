package welcomecontrollers

import (
	"encoding/json"
	"klikform/src/infras/configs"
	"net/http"
)

func WelcomeControllers(w http.ResponseWriter, r *http.Request) {
	// load informations from configs
	config := configs.LoadConfig()

	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Request success",
		"data": map[string]any{
			"about":   config.APP_NAME,
			"version": config.APP_VER,
		},
	})
}
