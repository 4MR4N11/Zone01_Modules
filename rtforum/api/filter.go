package api

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	c "forum/config"
	"forum/models"
	"forum/utils"
)

const (
	ALL = iota
	MY_POST
	LIKED_POST
)

func FilterApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		PostFilter(w, r)
	default:
		utils.WriteJSON(w, http.StatusMethodNotAllowed, "The HTTP method used in the request is invalid. Please ensure you're using the correct method.", nil)
	}
}

func getPaginationInfo(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	currPage, err := strconv.Atoi(pageStr)
	if err != nil || currPage < 1 {
		currPage = 1
	}
	return currPage, c.LIMIT_PER_PAGE
}

func PostFilter(w http.ResponseWriter, r *http.Request) {
	currPage, limit := getPaginationInfo(r)
	sessionID := utils.GetSessionCookie(r)
	session, err := c.SESSION.GetSession(sessionID)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	var userId int64 = -1
	if c.IsAuth(sessionID) != nil {
		userId = session.UserId
	}
	r.ParseForm()
	query := r.FormValue("query")
	options := r.FormValue("options")
	if options == "" {
		options = "0"
	}
	postType, err := strconv.Atoi(options)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "The filter criteria provided is not allowed", nil)
		return
	}
	postType, err = selectPostType(postType, userId != -1)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	postRep := models.NewPostRepository()

	posts, err := postRep.GetPostsBy(query, postType, userId, currPage, limit)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	posts, err = getPostsFilter(posts)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	count := len(posts)

	currentPagePosts := getCurrentPagePosts(posts, currPage, limit, count)
	page := NewPageStruct("forum", sessionID, nil)
	page.Data = IndexStruct{
		Posts:       currentPagePosts,
		TotalPages:  int(math.Ceil(float64(count) / c.LIMIT_PER_PAGE)),
		CurrentPage: currPage,
		Query:       query,
		Option:      postType,
	}
	utils.WriteJSON(w, http.StatusOK, "", page)
}

func getPostsFilter(posts []*models.Post) ([]*models.Post, error) {
	tagsRepo := models.NewTagRepository()

	for _, post := range posts {
		tags, err := tagsRepo.GetTagsForPost(post.ID)
		if err != nil {
			return nil, err
		}
		post.Tags = tags
	}
	return posts, nil
}

func getCurrentPagePosts(posts []*models.Post, currentPage int, limit int, count int) []*models.Post {
	if (currentPage-1)*limit > count {
		currentPage = max(int(math.Ceil(float64(count)/c.LIMIT_PER_PAGE)), 1)
	}
	return posts[(currentPage-1)*limit : min(count, (currentPage-1)*limit+limit)]
}

func selectPostType(value int, isAuth bool) (int, error) {
	if value < 0 || value > 2 {
		return 0, errors.New("the filter criteria provided is not allowed")
	}
	if isAuth {
		return value, nil
	}
	if value != 0 {
		return 0, errors.New("the selected filter is restricted to members only")
	}
	return 0, nil
}
