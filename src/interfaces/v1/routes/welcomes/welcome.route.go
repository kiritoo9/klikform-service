package welcomeroutes

import (
	welcomecontrollers "klikform/src/applications/controllers/welcomes"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

func WelcomeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/welcome", middlewares.Apply(
		welcomecontrollers.WelcomeControllers,
		middlewarecomponents.Method([]string{"GET"}),
	))
}
