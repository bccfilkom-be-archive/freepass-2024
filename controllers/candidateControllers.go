package controllers

import (
	"freepass-2024/initializers"
	"freepass-2024/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	user, _ := c.Get("user")
	if !user.(models.User).IsCandidate {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "you must be a candidate to create a post",
		})

		return
	}

	var body struct {
		Title string
		Body  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})

		return
	}

	post := models.Post{Author: user.(models.User).Username, Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create post",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created a post",
	})
}


func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	user, _ := c.Get("user")

	var post models.Post
	initializers.DB.First(&post, postID)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post not found",
		})

		return
	}

	if user.(models.User).Username != post.Author {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the post can only be updated by the author",
		})

		return
	}

	var body struct {
		NewTitle string
		NewBody  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})

		return
	}

	initializers.DB.Model(&post).Updates(models.Post{Title: body.NewTitle, Body: body.NewBody})

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully updated the post",
	})
}


func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	user, _ := c.Get("user")

	var post models.Post
	initializers.DB.First(postID)
	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post not found",
		})

		return
	}

	if user.(models.User).Username != post.Author {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the post can only be deleted by the author",
		})

		return
	}

	initializers.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted the post",
	})
}
