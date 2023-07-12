package posts

import (
	"blog/pkg/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatePost(service PostService) gin.HandlerFunc {
	return func(context *gin.Context) {
		postToCreate := &domain.Post{}
		err := context.Bind(postToCreate)
		if err != nil {
			context.Error(err)
		}

		postToCreate.AuthorID = context.GetUint("userId")
		postId, err := service.PublishPost(postToCreate)
		if err != nil {
			context.Error(err)
		}
		context.JSON(http.StatusCreated, gin.H{
			"message": "Post Created",
			"id":      postId,
		})
	}
}

func Index(service PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.Query("page"))
		posts, totalPages, err := service.GetPostByPage(page)

		if err != nil {
			c.Error(domain.NewAPIError(http.StatusBadRequest, "", err.Error()))
		}

		c.JSON(http.StatusOK, gin.H{
			"posts":      posts,
			"totalPages": totalPages,
		})
	}
}

func ReadPost(service PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, _ := strconv.Atoi(c.Param("id"))
		postFound, err := service.ReadPost(uint(postId))
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, postFound)
	}
}

func Reactions(service PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		postId, _ := strconv.Atoi(c.Param("id"))
		fmt.Println(postId)
		rt := c.Query("reaction")
		result, err := service.Reactions(uint(postId), rt)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

func CommentPost(service PostService) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
