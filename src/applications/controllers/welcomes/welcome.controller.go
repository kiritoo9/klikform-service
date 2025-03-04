package welcomecontrollers

import (
	"klikform/src/infras/configs"
	"klikform/src/utils"
	"net/http"
)

func WelcomeControllers(w http.ResponseWriter, r *http.Request) {
	// load informations from configs
	config := configs.LoadConfig()

	// set response
	utils.SetResponse(w, http.StatusBadRequest, "Request success", map[string]any{
		"about":   config.APP_NAME,
		"version": config.APP_VER,
	})
}
