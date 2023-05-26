package handlers

import (
	"encoding/json"
	"net/http"
	"user-go/models"
	"user-go/utils"
	"user-go/validations"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validations.ValidateUser(&user); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
