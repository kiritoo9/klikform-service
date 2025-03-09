package masterroutes

import (
	rolecontrollers "klikform/src/applications/controllers/masters"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	"net/http"
)

// @Summary      Role List
// @Description  List of role available
// @Tags         Master - Roles
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /roles [get]
func RoleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", middlewares.Apply(
		rolecontrollers.List,
		middlewarecomponents.Method("GET"),
		middlewarecomponents.Auth(),
	))
}
