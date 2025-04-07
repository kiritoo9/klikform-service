package authcontrollers

import (
	"encoding/json"
	"fmt"
	repos "klikform/src/applications/repos/masters"
	"klikform/src/infras/configs"
	schemas "klikform/src/interfaces/v1/schemas/auths"
	"klikform/src/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// @Summary      Authentication
// @Description  Get access token by login
// @Tags         Auth
// @Accept		 json
// @Param 		 request body schemas.AuthBodySchema true "Auth body"
// @Success      200  {object} schemas.ResponseSchema{data=schemas.AuthResponseSchema} "Successful response"
// @Failure		 400  {object} schemas.ResponseSchema "Failure response"
// @Router       /auth [post]
func Login(w http.ResponseWriter, r *http.Request) {
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
	user, err := repos.GetUserByEmail(body.Email)
	if err != nil {
		utils.SetResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Username and password does not match", nil)
		return
	}

	// generate token
	configs := configs.LoadConfig()
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(configs.JWT_SECRET)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Error while generating token", nil)
		return
	}

	// set response
	utils.SetResponse(w, http.StatusOK, "Auhtenticated", map[string]any{
		"access_token":  tokenString,
		"refresh_token": tokenString,
	})
}
