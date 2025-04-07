package masterroutes

import (
	mastercontrollers "klikform/src/applications/controllers/masters"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

// handle route based on method
func routeMethodHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mastercontrollers.UserList(w, r)
	case "POST":
		mastercontrollers.UserCreate(w, r)
	case "PUT":
		mastercontrollers.UserUpdate(w, r)
	case "DELETE":
		mastercontrollers.UserDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", middlewares.Apply(
		routeMethodHandlers,
		middlewarecomponents.Method([]string{"GET", "POST"}),
		middlewarecomponents.Auth(),
	))

	mux.HandleFunc("/users/", middlewares.Apply(
		routeMethodHandlers,
		middlewarecomponents.Method([]string{"GET", "DELETE", "PUT"}),
		middlewarecomponents.Auth(),
	))
}
