package api

import (
	"forum/config"
	"forum/services"
	"forum/utils"
	"net/http"
)

type AuthResponse struct {
    SessionID string `json:"session_id"`
    UserID    int    `json:"user_id"`
    Username  string `json:"username"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginApi(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := utils.ReadJSON(r, &creds)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "Invalid credentials", nil)
		return
	}
	user, err := services.LoginUser(creds.Password, creds.Username)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	session, err := config.SESSION.CreateSession(user.Username, user.ID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "An error occurred while creating a session", nil)
		return
	}
	cookies := http.Cookie{
		Name:    "session",
		Value:   session.ID,
		Expires: session.ExpiresAt,
		Path:    "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookies)
	response := AuthResponse{
		SessionID: session.ID,
		UserID:    int(user.ID),
		Username:  user.Username,
	}
	utils.WriteJSON(w, http.StatusOK, "Login successful", response)
}
