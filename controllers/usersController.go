package controllers

import (
	"freepass-2024/initializers"
	"freepass-2024/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	user := models.User{Username: body.Username, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username is already taken",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully created a user with the username " + body.Username,
	})
}

func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create JWT",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
}

func Edit(c *gin.Context) {
	var body struct {
		NewPassword string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read request body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	user, _ := c.Get("user")
	initializers.DB.Model(&user).Where("id = ?", user.(models.User).ID).Update("password", hash)
}

func FetchPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	if len(posts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch posts",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func ViewPost(c *gin.Context) {
	postID := c.Param("id")

	var post models.Post
	initializers.DB.Find(&post, postID)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to view post",
		})

		return
	}

	var comments []models.Comment
	initializers.DB.Where("post_id = ?", postID).Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"post":     post,
		"comments": comments,
	})
}

func FetchCandidates(c *gin.Context) {
	var candidates []models.User
	initializers.DB.Where("is_candidate = ?", true).Find(&candidates)

	if len(candidates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no candidates found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"candidates": candidates,
	})
}

func ViewCandidatesPosts(c *gin.Context) {
	var body struct {
		CandidateUsername string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var posts []models.Post
	initializers.DB.Where("author = ?", body.CandidateUsername).Find(&posts)

	if len(posts) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch posts",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func AddComment(c *gin.Context) {
	user, _ := c.Get("user")
	postID := c.Param("id")

	var body struct {
		Body string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var post models.Post
	initializers.DB.Find(&post, postID)

	if post.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to fetch post",
		})

		return
	}

	comment := models.Comment{PostID: post.ID, Username: user.(models.User).Username, Body: body.Body}
	result := initializers.DB.Create(&comment)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create comment",
		})

		return
	}
}

func CastVote(c *gin.Context) {
	user, _ := c.Get("user")

	if user.(models.User).HasVoted {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user has already voted",
		})

		return
	}

	var election models.Election
	initializers.DB.Last(&election)

	if !election.StartDate.Before(time.Now()) || !election.EndDate.After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "exceeded the election period",
		})

		return
	}

	candidateID := c.Param("id")

	var candidate models.User
	initializers.DB.Find(&candidate, candidateID)
	if candidate.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "candidate not found",
		})

		return
	}

	if !candidate.IsCandidate {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the selected user is not a candidate",
		})

		return
	}

	initializers.DB.Model(&candidate).Where("id = ?", candidateID).Update("votes", candidate.Votes+1)
	initializers.DB.Model(&user).Where("id = ?", user.(models.User).ID).Update("has_voted", true)
}
