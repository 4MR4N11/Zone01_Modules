package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/config"
	"forum/handlers"
	"forum/models"
	"forum/services"
	"forum/utils"
)

func PostApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePost(w, r)
	default:
		utils.WriteJSON(w, http.StatusMethodNotAllowed, "The HTTP method used in the request is invalid. Please ensure you're using the correct method.", nil)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	sessionId := utils.GetSessionCookie(r)
	session, _ := config.SESSION.GetSession(sessionId)
	if session == nil {
		utils.WriteJSON(w, http.StatusBadRequest, "You don't have the necessary permissions to access this. Please log in or check your access rights.", nil)
		return
	}
	var post models.Post
	err := utils.ReadJSON(r, &post)
	post.UserID = session.UserId
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if utils.IsEmpty(post.Content) || utils.IsEmpty(post.Title) {
		utils.WriteJSON(w, http.StatusBadRequest, "The title or content cannot be empty. Please provide both and try again.", nil)
		return
	}
	if len(post.Tags) > 5 {
		utils.WriteJSON(w, http.StatusBadRequest, "The tag count must be less than 6", nil)
		return
	}
	err = services.CreateNewPost(&post)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, "The post has been created successfully.", post)
}


type PostData struct {
	Post     models.Post
	Comments []models.Comment
}

func GetPostApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJSON(w, http.StatusUnauthorized, "The HTTP method used in the request is invalid. Please ensure you're using the correct method.", nil)
		return
	}
	postId, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	postRepo := models.NewPostRepository()
	comRepo := models.NewCommentRepository()
	post, err := postRepo.GetPostById(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.WriteJSON(w, http.StatusNotFound, "Not found", nil)
			return
		}
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	comment, err := comRepo.GetPostComments(postId)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	postData := PostData{
		Comments: comment,
		Post:     *post,
	}
	session := utils.GeTCookie("session", r)
	page := handlers.NewPageStruct(post.Title, session, postData)
	utils.WriteJSON(w, http.StatusOK, "OK", page)
}
