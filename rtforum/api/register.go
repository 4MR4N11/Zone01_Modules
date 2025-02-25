package api

import (
	"net/http"
	"strconv"
	"strings"

	"forum/models"
	"forum/services"
	"forum/utils"
)

func RegisterApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJSON(w, http.StatusMethodNotAllowed, "The HTTP method used in the request is invalid. Please ensure you're using the correct method.", nil)
		return
	}
	var user models.User
	err := utils.ReadJSON(r, &user)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	trimmedUsername := strings.TrimSpace(user.Username)
	user.Email = strings.ToLower(user.Email)
	if !utils.IsBetween(trimmedUsername, 3, 50) {
		utils.WriteJSON(w, http.StatusBadRequest, "The username must be between 3 and 50 characters in length", nil)
		return
	}
	if !utils.IsBetween(user.Password, 8, 50) {
		utils.WriteJSON(w, http.StatusBadRequest, "The password must be between 8 and 50 characters in length", nil)
		return
	}
	if !utils.IsValidEmail(user.Email) {
		utils.WriteJSON(w, http.StatusBadRequest, "The email address entered is invalid.", nil)
		return
	}
	if !utils.IsBetween(user.FirstName, 3, 25) {
		utils.WriteJSON(w, http.StatusBadRequest, "The full name must be between 3 and 25 characters in length", nil)
		return
	}
	if !utils.IsBetween(user.LastName, 3, 25) {
		utils.WriteJSON(w, http.StatusBadRequest, "The full name must be between 3 and 25 characters in length", nil)
		return
	}
	if user.Gender == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "Gender is required", nil)
		return
	}
	if user.Age == "" {
		user.Age = "0"
	}
	age, err := strconv.Atoi(user.Age)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "The age must be a numeric value", nil)
		return
	}
	if age < 16 {
		utils.WriteJSON(w, http.StatusBadRequest, "You must be at least 16 years old to register", nil)
		return
	}
	err = services.RegisterUser(&user)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "Registration completed successfully.", user)
}
