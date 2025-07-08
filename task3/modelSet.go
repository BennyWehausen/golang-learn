package task3

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	PostCount int `gorm:"default:0"` // 新增字段：用户的文章数量统计
	// 一对多关系：用户可以有多篇文章
	Posts []Post `gorm:"foreignKey:UserID"`
}

// Post 文章模型
type Post struct {
	ID      uint `gorm:"primaryKey"`
	Title   string
	Content string

	// 外键关联 User
	UserID uint `gorm:"index"` // 建立索引加快查询
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// 一对多关系：文章可以有多个评论
	Comments      []Comment `gorm:"foreignKey:PostID"` // ⚠️ 显式指定外键字段名
	CommentStatus int       `gorm:"default:0"`         // 新增字段：评论状态：0-有评论  1-无评论
}

// Comment 评论模型
type Comment struct {
	ID      uint `gorm:"primaryKey"`
	Content string

	// 外键关联 Post
	PostID uint
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
