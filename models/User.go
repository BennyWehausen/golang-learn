package models

import "gorm.io/gorm"

// 用户模型
type User struct {
	gorm.Model
	Username string    `gorm:"unique" json:"username"`
	Password string    `json:"password"`
	Email    string    `gorm:"unique" json:"email"`
	Posts    []Post    `json:"posts"`
	Comments []Comment `json:"comments"`
}
