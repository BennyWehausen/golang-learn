package task4

import (
	"github.com/gin-gonic/gin"
	"golang-learn/dao"
	"golang-learn/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// 用户注册
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查用户名是否已存在
	if exists := checkUsernameExists(user.Username); exists {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	user.Password = string(hashedPassword)

	if err := dao.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户注册成功"})
}

// 检查用户名是否存在
func checkUsernameExists(username string) bool {
	var user models.User
	result := dao.DB.Where("username = ?", username).First(&user)
	return result.Error == nil
}
