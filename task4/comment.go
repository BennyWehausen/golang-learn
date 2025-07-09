package task4

import (
	"github.com/gin-gonic/gin"
	"golang-learn/dao"
	"golang-learn/models"
	"net/http"
	"strconv"
)

// 创建评论
func CreateComment(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("currentUser")
	comment.UserID = user.(models.User).ID
	comment.PostID = uint(postID)

	if err := dao.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// 获取文章评论
func GetComments(c *gin.Context) {
	postIDParam := c.Param("id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var comments []models.Comment
	if err := dao.DB.Where("post_id = ?", postID).Preload("User").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取评论失败"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
