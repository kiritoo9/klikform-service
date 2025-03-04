package authroutes

import (
	authcontrollers "klikform/src/applications/controllers/auths"
	"klikform/src/interfaces/v1/middlewares"
	middlewarecomponents "klikform/src/interfaces/v1/middlewares/components"
	_ "klikform/src/interfaces/v1/schemas/auths" // call schema with unserscore(_) for swagger needs
	"net/http"
)

// @Summary      Authentication
// @Description  Get access token by login
// @Tags         Auth
// @Accept		 json
// @Param 		 request body schemas.AuthBodySchema true "Auth body"
// @Success      200  {object} schemas.ResponseSchema{data=schemas.AuthResponseSchema} "Successful response"
// @Failure		 400  {object} schemas.ResponseSchema "Failure response"
// @Router       /auth [post]
func AuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth", middlewares.Apply(
		authcontrollers.AuthController,
		middlewarecomponents.Method("POST"),
	))
}
