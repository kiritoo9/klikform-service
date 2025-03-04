package welcomeroutes

import (
	welcomecontrollers "klikform/src/applications/controllers/welcomes"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

// @Summary      Welcome point
// @Description  Welcome entry point to test API run
// @Tags         Welcome
// @Success      200  {object} schemas.ResponseSchema{data=schemas.WelcomeResponseSchema} "Successful response"
// @Router       /welcome [get]
func WelcomeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/welcome", middlewares.Apply(
		welcomecontrollers.WelcomeControllers,
		middlewarecomponents.Method("GET"),
	))
}
