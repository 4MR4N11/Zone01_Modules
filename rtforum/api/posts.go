package api

import (
	"forum/models"
	"forum/utils"
	"math"
	"net/http"
	"strconv"
	"forum/config"
)

type IndexStruct struct {
	Posts       []*models.Post
	TotalPages  int
	CurrentPage int
	Query       string
	Option      int
}

func GetPostsApi(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	currPage, err := strconv.Atoi(pageStr)
	if err != nil || currPage < 1 {
		currPage = 1
	}
	limit := config.LIMIT_PER_PAGE
	session := utils.GeTCookie("session", r)
	postRep := models.NewPostRepository()
	posts, err := getPosts(currPage, limit)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	count, err := postRep.Count()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	page := NewPageStruct("forum", session, nil)
	page.Data = IndexStruct{
		Posts:       posts,
		TotalPages:  int(math.Ceil(float64(count) / config.LIMIT_PER_PAGE)),
		CurrentPage: currPage,
	}
	utils.WriteJSON(w, http.StatusOK, "OK", page)
}

func getPosts(currPage, limit int) ([]*models.Post, error) {
	postRep := models.NewPostRepository()
	tagsRepo := models.NewTagRepository()
	posts, err := postRep.GetPostPerPage(currPage, limit)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		tags, err := tagsRepo.GetTagsForPost(post.ID)
		post.Content = post.Content[0:min(len(post.Content), 200)]
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	}
	return posts, nil
}
