package mastercontrollers

import (
	"encoding/json"
	"fmt"
	"klikform/src/applications/models"
	repos "klikform/src/applications/repos/masters"
	schemas "klikform/src/interfaces/v1/schemas/masters"
	"klikform/src/utils"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var userValidate = validator.New()

// @Summary      User List
// @Description  List of user available
// @Tags         Master - Users
// @Param        page   query     int     true  "Page number"
// @Param        limit  query     int     true  "Items per page"
// @Param        keywords  query  string  false  "Search keywords"
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /users [get]
func UserList(w http.ResponseWriter, r *http.Request) {
	// get parameters
	page := utils.AtoiOrDefault(r.URL.Query().Get("page"), 1)
	limit := utils.AtoiOrDefault(r.URL.Query().Get("limit"), 10)
	keywords := r.URL.Query().Get("keywords")

	// get data from database
	data, err := repos.GetUsers(page, limit, keywords)
	if err != nil {
		utils.SetResponse(w, http.StatusInternalServerError, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	totalPage := 0
	count, err := repos.GetCountUser(keywords)
	if err != nil {
		utils.SetResponse(w, http.StatusInternalServerError, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}
	if count != nil {
		totalCount := count.(int64)
		if totalCount > 0 {
			totalPage = int(totalCount) / limit
			if int(totalCount)%limit > 0 {
				totalPage++
			}
		}
	}

	// preparing data for response
	response := map[string]any{
		"page":       page,
		"total_page": totalPage,
		"limit":      limit,
		"keywords":   keywords,
		"data":       data,
	}

	// send response
	utils.SetResponse(w, http.StatusOK, "Request success", response)
	return
}

// @Summary      User Detail
// @Description  Detail of user
// @Tags         Master - Users
// @Param        id  path  string  true  "ID of user"
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /users/{id} [get]
func UserDetail(w http.ResponseWriter, r *http.Request) {
	// get parameters
	params := strings.TrimPrefix(r.URL.Path, "/users/")
	paths := strings.Split(params, "/")
	id := paths[0]

	if id == "" {
		utils.SetResponse(w, http.StatusBadRequest, "ID is required", nil)
		return
	}

	// get data from database
	user, err := repos.GetUserById(id)
	if err != nil {
		utils.SetResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}
	userRole, _ := repos.GetRoleByUser(id) // leave error handling

	// prepare to response
	response := map[string]any{
		"user": user,
		"user_role": map[string]any{
			"role_id": userRole.RoleID,
			"role":    userRole.Role.Name,
		},
	}

	// send response
	utils.SetResponse(w, http.StatusOK, "Request success", response)
	return
}

// @Summary      User Create
// @Description  Create new user
// @Tags         Master - Users
// @Accept		 json
// @Param 		 request body schemas.UserBodySchema true "User body"
// @Success      200  {object} schemas.ResponseSchema{data=schemas.UserResponseSchema} "Successful response"
// @Failure		 400  {object} schemas.ResponseSchema "Failure response"
// @Router		 /users [post]
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var body schemas.UserBodySchema
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": "Invalid JSON Body",
		})
		return
	}

	err = userValidate.Struct(body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Error validation body", map[string]any{
			"error": fmt.Sprintf("Validation error: %v", err),
		})
		return
	}

	// check available email
	_, err = repos.GetUserByEmail(body.Email)
	if err == nil {
		utils.SetResponse(w, http.StatusBadRequest, "Email is already exists", nil)
		return
	}

	// preparing data and perform to insert
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something wrong went validating your password", nil)
		return
	}

	user := models.Users{
		ID:       uuid.New(),
		Email:    body.Email,
		Password: string(hashedBytes),
		Fullname: body.Fullname,
		Phone:    body.Phone,
		Address:  body.Address,
		Status:   body.Status,
	}

	roleID, err := uuid.Parse(body.RoleID)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Invalid role ID", nil)
		return
	}
	userRole := models.UserRoles{
		ID:     uuid.New(),
		UserID: user.ID,
		RoleID: roleID,
	}

	result, err := repos.CreateUser(user, userRole)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	// send response
	utils.SetResponse(w, http.StatusCreated, "Data inserted", map[string]any{
		"id":         result.ID,
		"email":      result.Email,
		"fullname":   result.Fullname,
		"phone":      result.Phone,
		"address":    result.Address,
		"status":     result.Status,
		"created_at": result.CreatedAt,
	})
	return
}

// @Summary      User Update
// @Description  Update existing user
// @Tags         Master - Users
// @Accept		 json
// @Param        id  path  string  true  "ID of user"
// @Param 		 request body schemas.UserBodySchema true "User body"
// @Success      204 "Successful response"
// @Failure		 400  {object} schemas.ResponseSchema "Failure response"
// @Router		 /users/{id} [put]
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	var body schemas.UserBodySchema
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": "Invalid JSON Body",
		})
		return
	}

	err = userValidate.Var(body.Password, "omitempty")
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Error validation body", map[string]any{
			"error": fmt.Sprintf("Validation error: %v", err),
		})
		return
	}

	// get parameters
	params := strings.TrimPrefix(r.URL.Path, "/users/")
	paths := strings.Split(params, "/")
	id := paths[0]

	if id == "" {
		utils.SetResponse(w, http.StatusBadRequest, "ID is required", nil)
		return
	}

	// check existing user
	user, err := repos.GetUserById(id)
	if err != nil {
		utils.SetResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// check existing user role
	existingUserRole, _ := repos.GetRoleByUser(id)

	// check available email
	_, err = repos.GetUserByEmail(body.Email, id)
	if err == nil {
		utils.SetResponse(w, http.StatusBadRequest, "Email is already exists", nil)
		return
	}

	// preparing data to update
	user.Email = body.Email
	user.Fullname = body.Fullname
	user.Phone = body.Phone
	user.Address = body.Address
	user.Status = body.Status
	user.UpdatedAt = time.Now()

	// generate password if exists
	if body.Password != "" {
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
		if err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "Something wrong went validating your password", nil)
			return
		}
		user.Password = string(hashedBytes)
	}

	roleID, err := uuid.Parse(body.RoleID)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Invalid role ID", nil)
		return
	}

	var userRole map[string]any
	if existingUserRole == nil {
		userRole = map[string]any{
			"id":      uuid.New(),
			"user_id": user.ID,
			"role_id": roleID,
		}
	} else {
		userRole = map[string]any{
			"role_id":    roleID,
			"updated_at": time.Now(),
		}
	}

	_, err = repos.UpdateUser(user, userRole)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	// set response
	utils.SetResponse(w, http.StatusNoContent, "Data updated", nil)
	return
}

// @Summary      User Delete
// @Description  Delete existing user
// @Tags         Master - Users
// @Param        id  path  string  true  "ID of user"
// @Success      204 "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /users/{id} [delete]
func UserDelete(w http.ResponseWriter, r *http.Request) {
	// get parameters
	params := strings.TrimPrefix(r.URL.Path, "/users/")
	paths := strings.Split(params, "/")
	id := paths[0]

	if id == "" {
		utils.SetResponse(w, http.StatusBadRequest, "ID is required", nil)
		return
	}

	// check existing user
	user, err := repos.GetUserById(id)
	if err != nil {
		utils.SetResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// check existing user role
	existingUserRole, _ := repos.GetRoleByUser(id)

	// preparing data to delete
	user.Deleted = true
	user.UpdatedAt = time.Now()

	var userRole map[string]any
	if existingUserRole != nil {
		userRole = map[string]any{
			"deleted":    true,
			"updated_at": time.Now(),
		}
	}

	// perform to delete
	_, err = repos.UpdateUser(user, userRole)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	// send response
	utils.SetResponse(w, http.StatusNoContent, "Data deleted", nil)
	return
}
