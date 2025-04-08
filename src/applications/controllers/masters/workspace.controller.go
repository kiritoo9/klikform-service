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

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var workspaceValidate = validator.New()

// @Summary      Workspace List
// @Description  List of workspace available
// @Tags         Master - Workspaces
// @Param        page   query     int     true  "Page number"
// @Param        limit  query     int     true  "Items per page"
// @Param        keywords  query  string  false  "Search keywords"
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /workspaces [get]
func WorkspaceList(w http.ResponseWriter, r *http.Request) {
	// get logged data
	token := r.Context().Value("loggedToken")
	var userID string
	role := token.(map[string]any)["role"].(string)
	if strings.ToLower(role) == "admin" {
		userID = token.(map[string]any)["id"].(string)
	} else {
		userID = ""
	}

	// get parameters
	page := utils.AtoiOrDefault(r.URL.Query().Get("page"), 1)
	limit := utils.AtoiOrDefault(r.URL.Query().Get("limit"), 10)
	keywords := r.URL.Query().Get("keywords")

	// get data from database
	data, err := repos.GetWorkspaces(page, limit, keywords, userID)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	totalPage := 0
	count, err := repos.GetCountWorkspace(keywords, userID)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
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

	// send response
	utils.SetResponse(w, http.StatusOK, "Request success", map[string]any{
		"page":       page,
		"total_page": totalPage,
		"limit":      limit,
		"keywords":   keywords,
		"data":       data,
	})
	return
}

// @Summary      Workspace Detail
// @Description  Detail of workspace
// @Tags         Master - Workspaces
// @Param        id  path  string  true  "ID of workspace"
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /workspaces/{id} [get]
func WorkspaceDetail(w http.ResponseWriter, r *http.Request) {
	// get parameters
	params := strings.TrimPrefix(r.URL.Path, "/workspaces/")
	paths := strings.Split(params, "/")
	id := paths[0]
	if id == "" {
		utils.SetResponse(w, http.StatusBadRequest, "ID is required", nil)
		return
	}

	// get detail from database
	data, err := repos.GetWorkspaceById(id)
	if err != nil {
		utils.SetResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	// send response
	response := map[string]any{
		"data": data,
	}
	utils.SetResponse(w, http.StatusOK, "Request success", response)
	return
}

// @Summary      Workspace Create
// @Description  Create new workspace
// @Tags         Master - Workspaces
// @Accept		 json
// @Param 		 request body schemas.WorkspaceBodySchema true "Workspace body"
// @Success      200  {object} schemas.ResponseSchema{data=schemas.WorkspaceResponseSchema} "Successful response"
// @Failure		 400  {object} schemas.ResponseSchema "Failure response"
// @Router		 /workspaces [post]
func WorkspaceCreate(w http.ResponseWriter, r *http.Request) {
	// get logged data
	token := r.Context().Value("loggedToken")
	userID := token.(map[string]any)["id"].(string)
	if userID == "" {
		utils.SetResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var body schemas.WorkspaceBodySchema
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": "Invalid JSON Body",
		})
		return
	}

	err = workspaceValidate.Struct(body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Error validation body", map[string]any{
			"error": fmt.Sprintf("Validation error: %v", err),
		})
		return
	}

	// prepare and perform to insert data
	workspace := models.Workspaces{
		ID:           uuid.New(),
		Title:        body.Title,
		Descriptions: body.Descriptions,
		Status:       body.Status,
		Remark:       body.Remark,
	}

	workspaceUser := models.WorkspaceUsers{
		ID:            uuid.New(),
		WorkspaceID:   workspace.ID,
		UserID:        userUUID,
		IsOwner:       true,
		AccessControl: "A2",
		Remark:        "owner",
	}

	data, err := repos.CreateWorkspace(workspace, workspaceUser)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	utils.SetResponse(w, http.StatusCreated, "Request success", map[string]any{
		"data": data,
	})
	return
}

func WorkspaceUpdate(w http.ResponseWriter, r *http.Request) {
	utils.SetResponse(w, http.StatusNoContent, "Request success", nil)
	return
}

func WorkspaceDelete(w http.ResponseWriter, r *http.Request) {
	utils.SetResponse(w, http.StatusNoContent, "Request success", nil)
	return
}
