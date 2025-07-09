package task4

import (
	"github.com/gin-gonic/gin"
	"golang-learn/dao"
	"golang-learn/models"
	"net/http"
	"strconv"
)

// 创建文章
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("currentUser")
	post.UserID = user.(models.User).ID

	if err := dao.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// 获取所有文章
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := dao.DB.Preload("User").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// 获取单个文章
func GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var post models.Post
	if err := dao.DB.Preload("User").Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 更新文章
func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var post models.Post
	if err := dao.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	user, _ := c.Get("currentUser")
	if post.UserID != user.(models.User).ID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无权修改该文章"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dao.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新文章失败"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// 删除文章
func DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var post models.Post
	if err := dao.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	user, _ := c.Get("currentUser")
	if post.UserID != user.(models.User).ID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无权删除该文章"})
		return
	}

	if err := dao.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除文章失败"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
