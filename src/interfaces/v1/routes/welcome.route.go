package routes

import (
	welcomecontrollers "klikform/src/applications/controllers"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

func WelcomeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middlewares.Apply(
		welcomecontrollers.WelcomeControllers,
		middlewarecomponents.Method("GET"),
	))
}
