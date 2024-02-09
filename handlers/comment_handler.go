package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func Comment(c *gin.Context) {
	// Mendapatkan ID pengguna dari token
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token tidak valid"})
		return
	}

	// Membaca data komentar dari permintaan
	var newComment models.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menetapkan UserID ke ID pengguna yang ditemukan dari token
	newComment.UserID = userID
	newComment.ID = len(models.Comments) + 1
	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pos tidak valid"})
		return
	}

	newComment.PostID = postID

	// Mencari postingan berdasarkan PostID
	var foundPost bool
	for _, post := range models.Posts {
		if post.ID == newComment.PostID {
			models.Comments = append(models.Comments, newComment)
			foundPost = true
			break
		}
	}

	if !foundPost {
		c.JSON(http.StatusNotFound, gin.H{"error": "Postingan tidak ditemukan"})
		return
	}

	// Setelah komentar ditambahkan, Anda dapat mengirimkan respons yang sesuai
	c.JSON(http.StatusOK, gin.H{"message": "Komentar berhasil ditambahkan"})
}

func DeleteComment(c *gin.Context) {
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

	commentID, err := strconv.Atoi(c.Param("commentID")) // Ubah ke "commentID"
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID komentar tidak valid"})
		return
	}

	for i, comment := range models.Comments {
		if comment.ID == commentID {
			models.Comments = append(models.Comments[:i], models.Comments[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Komentar berhasil dihapus"})
			return
		}
	}

	// Jika postingan tidak ditemukan
	c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
}
