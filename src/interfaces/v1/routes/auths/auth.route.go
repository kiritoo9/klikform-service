package authroutes

import (
	authcontrollers "klikform/src/applications/controllers/auths"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	_ "klikform/src/interfaces/v1/schemas/auths" // call schema with unserscore(_) for swagger needs
	"net/http"
)

func AuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth", middlewares.Apply(
		authcontrollers.Login,
		middlewarecomponents.Method([]string{"POST"}),
	))
}
