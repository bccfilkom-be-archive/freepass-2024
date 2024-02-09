package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post

	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Periksa apakah title dan description kosong
	if newPost.Title == "" || newPost.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and description cannot be empty"})
		return
	}
	if len(models.Settings) == 0 || time.Now().Before(models.Settings[0].Dates) {

		// Set timestamp
		newPost.Dates = time.Now()

		// Inisialisasi ID (contoh sederhana, Anda mungkin membutuhkan mekanisme lain untuk ID)
		newPost.ID = len(models.Posts) + 1

		// Tambahkan postingan baru ke slice posts
		models.Posts = append(models.Posts, newPost)

		c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": newPost})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Waktu posting telah berakhir"})
	}
}

func ViewPosts(c *gin.Context) {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	// Buat map untuk menyimpan komentar berdasarkan ID posnya
	commentMap := make(map[int][]models.Comment)
	for _, comment := range models.Comments {
		commentMap[comment.PostID] = append(commentMap[comment.PostID], comment)
	}

	// Buat slice untuk menyimpan posting dengan komentar
	var postsWithComments []gin.H
	for _, post := range models.Posts {
		postWithComments := gin.H{
			"id":          post.ID,
			"title":       post.Title,
			"description": post.Description,
			"comments":    commentMap[post.ID], // Ambil komentar yang sesuai dengan ID posting
		}
		postsWithComments = append(postsWithComments, postWithComments)
	}

	c.JSON(http.StatusOK, gin.H{"posts": postsWithComments})
}

func UpdatePost(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))

	var updatedPost *models.Post
	for i, post := range models.Posts {
		if post.ID == postID {
			updatedPost = &models.Posts[i]
			break
		}
	}

	if updatedPost == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Memeriksa dan memperbarui judul dan deskripsi
	var updateData struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update judul dan deskripsi postingan
	updatedPost.Title = updateData.Title
	updatedPost.Description = updateData.Description

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))

	// Temukan postingan berdasarkan ID dan hapus
	for i, post := range models.Posts {
		if post.ID == postID {
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
			return
		}
	}

	// Jika postingan tidak ditemukan
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}
