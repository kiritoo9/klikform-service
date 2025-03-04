package authcontrollers

import (
	"encoding/json"
	"fmt"
	schemas "klikform/src/interfaces/v1/schemas/auths"
	"klikform/src/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func AuthController(w http.ResponseWriter, r *http.Request) {
	// validate request-body
	var body schemas.AuthBodySchema
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": "Invalid JSON body",
		})
		return
	}

	err = validate.Struct(body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Error validation body", map[string]any{
			"error": fmt.Sprintf("Validation error: %v", err),
		})
		return
	}

	// check existing email

	// verify password

	// generate token

	// set response
	utils.SetResponse(w, http.StatusOK, "Auhtenticated", map[string]any{
		"access_token":  nil,
		"refresh_token": nil,
	})
}
