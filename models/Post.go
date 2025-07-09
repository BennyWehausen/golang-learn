package models

import "gorm.io/gorm"

// 博客文章模型
type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"comments"`
}
