package mastercontrollers

import (
	repos "klikform/src/applications/repos/masters"
	"klikform/src/utils"
	"net/http"
)

// @Summary      Role List
// @Description  List of role available
// @Tags         Master - Roles
// @Param        page   query     int     true  "Page number"
// @Param        limit  query     int     true  "Items per page"
// @Param        keywords  query  string  false  "Search keywords"
// @Success      200 {object} map[string]interface{} "Successful response"
// @Failure      400 {object} map[string]interface{} "Failure response"
// @Router		 /roles [get]
func RoleList(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters
	page := utils.AtoiOrDefault(r.URL.Query().Get("page"), 1)
	limit := utils.AtoiOrDefault(r.URL.Query().Get("limit"), 10)
	keywords := r.URL.Query().Get("keywords")

	// call repo to get the data
	role, err := repos.GetRoles(page, limit, keywords)
	if err != nil {
		utils.SetResponse(w, http.StatusInternalServerError, "Something went wrong", map[string]any{
			"error": err.Error(),
		})
		return
	}

	totalPage := 0
	count, err := repos.GetCountRoles(keywords)
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

	// preparet data for response
	response := map[string]any{
		"page":       page,
		"total_page": totalPage,
		"limit":      limit,
		"keywords":   keywords,
		"data":       role,
	}

	// send response
	utils.SetResponse(w, http.StatusOK, "Request success", response)
	return
}
