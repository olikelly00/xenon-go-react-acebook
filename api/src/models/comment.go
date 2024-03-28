package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `json:"message"`
	Likes   int    `json:"likes" gorm:"default=0"`
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
}

func (comment *Comment) Save() (*Comment, error) {
	err := Database.Create(comment).Error
	if err != nil {
		return &Comment{}, err
	}

	return comment, nil
}

// below function is different from posts as fetching all comments for a single post
// unsure what data type argument should be, was thinking int8 as per table structure
// but Users FindUserById function uses string instead
func FetchAllCommentsByPostId(post_id string) (*[]Comment, error) {
	var comments []Comment
	err := Database.Where("post_id = ?", post_id).First(&comments).Error

	fmt.Println(comments)

	if err != nil {
		return &[]Comment{}, err
	}

	return &comments, nil
}
