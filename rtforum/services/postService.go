package services

import (
	"errors"
	"strings"
	"time"

	"forum/models"
	"forum/utils"
)

func CheckTags(strs []string) bool {
	for _, str := range strs {
		if len(str) > 20 {
			return false
		}
	}
	return true
}

func CreateNewPost(post *models.Post) error {
	postRepo := models.NewPostRepository()
	TagsRepo := models.NewTagRepository()
	if !utils.IsBetween(post.Title, 0, 200) {
		return errors.New("title has exceeded the limits")
	}
	if !utils.IsBetween(post.Content, 0, 3000) {
		return errors.New("content has exceeded the limits")
	}
	if !CheckTags(post.Tags) {
		return errors.New("tags have exceeded the limits")
	}
	// check if input empty
	if strings.TrimSpace(post.Content) == "" || post.IsTagsEmpty() {
		return errFieldsEmpty
	}
	post.CreatedAt = time.Now()
	err := postRepo.Create(post)
	if err != nil {
		return err
	}
	TagsRepo.LinkTagsToPost(post.ID, post.Tags)
	return err
}
