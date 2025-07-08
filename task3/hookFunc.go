package task3

import "gorm.io/gorm"

/**
 * 1. 在 Post 模型中定义一个 AfterCreate 钩子函数，用于在创建文章后更新用户文章数量
 */
func (post *Post) AfterCreate(tx *gorm.DB) error {
	var user User
	if err := tx.First(&user, post.UserID).Error; err != nil {
		return err
	}

	// 更新用户文章数量 +1
	user.PostCount++
	return tx.Save(&user).Error
}

/**
 * 2. 在 Comment 模型中定义一个 AfterDelete 钩子函数，用于在删除评论后更新文章的评论状态
 */
func (comment *Comment) AfterDelete(tx *gorm.DB) error {
	var count int64
	var post Post

	// 获取该评论所属文章
	if err := tx.First(&post, comment.PostID).Error; err != nil {
		return err
	}

	// 查询当前文章下的评论总数
	tx.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&count)

	// 如果评论数量为 0，则更新文章的评论状态
	if count == 0 {
		post.CommentStatus = 1
		return tx.Save(&post).Error
	}
	return nil
}
