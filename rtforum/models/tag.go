package models

import (
	"database/sql"

	"forum/config"
)

type Tag struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TagRepository struct {
	db *sql.DB
}

func NewTagRepository() *TagRepository {
	return &TagRepository{db: config.DB}
}

func (r *TagRepository) CreateTag(name string) (*Tag, error) {
	query := "INSERT INTO tags (name) VALUES (?)"
	res, err := r.db.Exec(query, name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &Tag{ID: id, Name: name}, nil
}

func (r *TagRepository) GetTagsForPost(postId int64) ([]string, error) {
	query := `SELECT t.name
  FROM tags t
	JOIN post_tags pt ON t.id = pt.tagId
  WHERE pt.postId = ?`

	rows, err := r.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tagName string
		if err := rows.Scan(&tagName); err != nil {
			if err == sql.ErrNoRows {
				return tags, nil
			}
			return nil, err
		}
		tags = append(tags, tagName)
	}

	return tags, nil
}

func (r *TagRepository) LinkTagsToPost(postId int64, tagNames []string) error {
	selectStmt, err := r.db.Prepare("SELECT id FROM tags WHERE name = ?")
	if err != nil {
		return err
	}
	defer selectStmt.Close()

	insertStmt, err := r.db.Prepare("INSERT INTO post_tags (postId, tagId) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	for _, tagName := range tagNames {
		var tagId int64

		err := selectStmt.QueryRow(tagName).Scan(&tagId)
		if err != nil {
			if err == sql.ErrNoRows {
				newTag, err := r.CreateTag(tagName)
				if err != nil {
					return err
				}
				tagId = newTag.ID
			} else {
				return err
			}
		}

		_, err = insertStmt.Exec(postId, tagId)
		if err != nil {
			return err
		}
	}

	return nil
}
