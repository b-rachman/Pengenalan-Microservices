package handler

import (
	"net/http"

	"github.com/b-rachman/Pengenalan-Microservices/utils"
)

func ValidateAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("Authorization")
	if token != "asdfghjk" {
		utils.WrapAPIError(w, r, "Auth Invalid", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}
