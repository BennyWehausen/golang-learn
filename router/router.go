package router

import (
	"github.com/gin-gonic/gin"
	"golang-learn/task4"
	"golang-learn/utils"
)

func Router(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", task4.Register)
			auth.POST("/login", task4.Login)
		}

		posts := api.Group("/posts")
		posts.Use(utils.AuthMiddleware())
		{
			posts.POST("/", task4.CreatePost)
			posts.GET("/", task4.GetPosts)
			posts.GET("/:id", task4.GetPost)
			posts.PUT("/:id", task4.UpdatePost)
			posts.DELETE("/:id", task4.DeletePost)

			comments := posts.Group("/:id/comments")
			{
				comments.POST("/", task4.CreateComment)
				comments.GET("/", task4.GetComments)
			}
		}
	}
	return r
}
