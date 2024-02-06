package controllers

import (
	"freepass-2024/initializers"
	"freepass-2024/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PromoteUser(c *gin.Context) {
	userAdmin, _ := c.Get("user")
	if userAdmin.(models.User).Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not an admin",
		})

		return
	}

	userID := c.Param("id")
	var user models.User
	initializers.DB.Find(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})

		return
	}

	initializers.DB.Model(&user).Where("id = ?", userID).Update("is_candidate", true)
}

func DeleteUser(c *gin.Context) {
	userAdmin, _ := c.Get("user")
	if userAdmin.(models.User).Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not an admin",
		})

		return
	}

	userID := c.Param("id")
	var user models.User
	initializers.DB.Find(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not found",
		})

		return
	}

	initializers.DB.Delete(&user)
}

func DeletePostAdmin(c *gin.Context) {
	userAdmin, _ := c.Get("user")
	if userAdmin.(models.User).Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not an admin",
		})

		return
	}

	postID := c.Param("id")
	var post models.Post
	initializers.DB.Find(&post, postID)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "post not found",
		})

		return
	}

	initializers.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted the post",
	})
}

func DeleteComment(c *gin.Context) {
	userAdmin, _ := c.Get("user")
	if userAdmin.(models.User).Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not an admin",
		})

		return
	}

	commentID := c.Param("id")
	var comment models.Comment
	initializers.DB.Find(&comment, commentID)

	if comment.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "comment not found",
		})

		return
	}

	initializers.DB.Delete(&comment)
}

func SetElection(c *gin.Context) {
	userAdmin, _ := c.Get("user")
	if userAdmin.(models.User).Username != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you are not an admin",
		})

		return
	}

	var body struct {
		Year  string
		Start string
		End   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})

		return
	}

	const layout = "2006-01-02"
	parsedStart, err1 := time.Parse(layout, body.Start)
	parsedEnd, err2 := time.Parse(layout, body.End)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed parsing date",
			"a":     err1,
			"b":     err2,
		})

		return
	}
	
	election := models.Election{Year: uint(time.Now().Year()), StartDate: parsedStart, EndDate: parsedEnd}
	result := initializers.DB.Create(&election)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to set election period",
		})

		return
	}
}
