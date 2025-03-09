package rolecontrollers

import (
	"klikform/src/utils"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	utils.SetResponse(w, http.StatusOK, "Request success", nil)
	return
}
