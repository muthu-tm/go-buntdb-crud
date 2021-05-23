package controllers

import (
	"encoding/json"
	"go-inmem-crud/db"
	"go-inmem-crud/utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmp(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := db.Employee{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.AddEmployee(db.GetDB())

	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusCreated, userCreated)
}

func GetAllEmp(w http.ResponseWriter, r *http.Request) {

	user := db.Employee{}

	users, err := user.FindAllEmployees(db.GetDB())
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, users)
}

func GetEmp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userGotten, err := db.FindByID(db.GetDB(), vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, userGotten)
}

func UpdateEmp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := db.Employee{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("update")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedUser, err := user.UpdateAUser(db.GetDB(), vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, updatedUser)
}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	err := db.DeleteEmp(db.GetDB(), vars["id"])
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Removed an Item successfully!"})
}
