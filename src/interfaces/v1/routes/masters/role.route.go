package masterroutes

import (
	mastercontrollers "klikform/src/applications/controllers/masters"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

func RoleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/roles", middlewares.Apply(
		mastercontrollers.RoleList,
		middlewarecomponents.Method([]string{"GET"}),
		middlewarecomponents.Auth(),
	))
}
