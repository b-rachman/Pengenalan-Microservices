package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/b-rachman/Pengenalan-Microservices/menu-service/database"
	"github.com/b-rachman/Pengenalan-Microservices/utils"
	"gorm.io/gorm"
)

type MenuHandler struct {
	Db *gorm.DB
}

func (handler *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	var menu database.Menu
	err = json.Unmarshal(body, &menu)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = menu.Insert(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(w, r, "success", http.StatusOK)
}

func (handler *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusInternalServerError)
		return
	}

	menu := database.Menu{}
	menus, err := menu.GetAll(handler.Db)
	if err != nil {
		utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, http.StatusOK, "success")
}
