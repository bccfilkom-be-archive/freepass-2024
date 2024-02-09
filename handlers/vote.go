package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func Vote(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token tidak valid"})
		return
	}

	// Pemeriksaan apakah pengguna sudah memberikan vote sebelumnya
	for _, v := range models.Votes {
		if v.UserID == userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah memberikan vote sebelumnya"})
			return
		}
	}

	// Mendapatkan candidateID dari URL
	candidateID, err := strconv.Atoi(c.Param("candidateID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID kandidat tidak valid"})
		return
	}

	// Membuat objek vote baru
	var newVote models.Vote
	newVote.ID = len(models.Votes) + 1
	newVote.UserID = userID
	newVote.CandidateID = candidateID

	// Menyimpan vote ke slice votes
	models.Votes = append(models.Votes, newVote)

	c.JSON(http.StatusOK, gin.H{"message": "Vote casted successfully", "vote": newVote})
}

func VoteTotal(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token tidak valid"})
		return
	}

	// Membuat map untuk menyimpan jumlah vote setiap kandidat
	voteCount := make(map[int]int)

	// Menghitung jumlah vote setiap kandidat
	for _, v := range models.Votes {
		voteCount[v.CandidateID]++
	}

	// Membuat slice untuk menyimpan hasil
	var results []gin.H

	// Mengonversi jumlah vote menjadi bentuk yang diinginkan dalam response
	for candidateID, count := range voteCount {
		// Mencari nama kandidat berdasarkan ID kandidat
		var candidateName string
		for _, user := range models.Users {
			if user.ID == candidateID {
				candidateName = user.Name
				break
			}
		}

		// Menambahkan data kandidat dan total vote ke slice hasil
		result := gin.H{
			"ID":    candidateID,
			"Name":  candidateName,
			"Total": count,
		}
		results = append(results, result)
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
