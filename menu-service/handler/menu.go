package handler

import (
	"net/http"

	"github.com/FadhlanHawali/Digitalent-Kominfo_Pendalaman-Rest-API/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}
