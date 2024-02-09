package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func createPost(c *gin.Context) {
	var newPost Post

	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	// Pemeriksaan peran
	if getUserRoleFromToken(c) != "candidate" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only candidate can create posts"})
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
	if len(settings) == 0 || time.Now().Before(settings[0].Dates) {

		// Set timestamp
		newPost.Dates = time.Now()

		// Inisialisasi ID (contoh sederhana, Anda mungkin membutuhkan mekanisme lain untuk ID)
		newPost.ID = len(posts) + 1

		// Tambahkan postingan baru ke slice posts
		posts = append(posts, newPost)

		c.JSON(http.StatusOK, gin.H{"message": "Post created successfully", "post": newPost})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Waktu posting telah berakhir"})
	}
}

func viewPosts(c *gin.Context) {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	// Buat map untuk menyimpan komentar berdasarkan ID posnya
	commentMap := make(map[int][]Comment)
	for _, comment := range comments {
		commentMap[comment.PostID] = append(commentMap[comment.PostID], comment)
	}

	// Buat slice untuk menyimpan posting dengan komentar
	var postsWithComments []gin.H
	for _, post := range posts {
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

func updatePost(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	// Pemeriksaan peran admin
	role := getUserRoleFromToken(c)
	if role != "admin" && role != "candidate" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only admin or candidate can update posts"})
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))

	var updatedPost *Post
	for i, post := range posts {
		if post.ID == postID {
			updatedPost = &posts[i]
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

func deletePost(c *gin.Context) {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token tidak valid"})
		return
	}

	// Pemeriksaan peran admin
	role := getUserRoleFromToken(c)
	if role != "admin" && role != "candidate" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only admin or candidate can update posts"})
		return
	}

	postID, err := strconv.Atoi(c.Param("postID"))

	// Temukan postingan berdasarkan ID dan hapus
	for i, post := range posts {
		if post.ID == postID {
			posts = append(posts[:i], posts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
			return
		}
	}

	// Jika postingan tidak ditemukan
	c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
}
