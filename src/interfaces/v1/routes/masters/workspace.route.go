package masterroutes

import (
	mastercontrollers "klikform/src/applications/controllers/masters"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

func workspaceRouteMethodHandlers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/workspaces" {
			mastercontrollers.WorkspaceList(w, r)
		} else {
			mastercontrollers.WorkspaceDetail(w, r)
		}
	case "POST":
		mastercontrollers.WorkspaceCreate(w, r)
	case "PUT":
		mastercontrollers.WorkspaceUpdate(w, r)
	case "DELETE":
		mastercontrollers.WorkspaceDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func WorkspaceRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/workspaces", middlewares.Apply(
		workspaceRouteMethodHandlers,
		middlewarecomponents.Method([]string{"GET", "POST"}),
		middlewarecomponents.Auth(),
	))
	mux.HandleFunc("/workspaces/", middlewares.Apply(
		workspaceRouteMethodHandlers,
		middlewarecomponents.Method([]string{"GET", "PUT", "DELETE"}),
		middlewarecomponents.Auth(),
	))
}
