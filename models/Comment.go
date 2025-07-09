package models

import "gorm.io/gorm"

// 评论模型
type Comment struct {
	gorm.Model
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
	User    User   `json:"user"`
	PostID  uint   `json:"post_id"`
	Post    Post   `json:"post"`
}
