package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func setElectionDates(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token tidak valid"})
		return
	}

	// Pemeriksaan peran admin
	if getUserRoleFromToken(c) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only admin can delete users"})
		return
	}

	var newSetting Setting
	if err := c.ShouldBindJSON(&newSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada pengaturan tanggal pemilihan, tambahkan satu
	if len(settings) == 0 {
		newSetting.ID = 1
		settings = append(settings, newSetting)
		c.JSON(http.StatusOK, gin.H{"message": "Election dates set successfully", "newSetting": newSetting})
		return
	}

	// Jika ada pengaturan tanggal pemilihan, ganti dengan yang baru
	settings[0] = newSetting
	c.JSON(http.StatusOK, gin.H{"message": "Election dates updated successfully", "newSetting": newSetting})
}
func promoteToCandidate(c *gin.Context) {
	// Pemeriksaan otentikasi
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token tidak valid"})
		return
	}

	// Pemeriksaan peran admin
	if getUserRoleFromToken(c) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only admin can delete users"})
		return
	}

	// Mendapatkan candidateID dari URL
	candidateID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	var updatedUser User
	userFound := false
	// Mencari pengguna dengan ID yang sesuai dan memperbarui perannya menjadi "candidate"
	for i, u := range users {
		if u.ID == candidateID {
			users[i].Role = "candidate"
			users[i].Token = ""
			updatedUser = users[i]
			userFound = true
			break
		}
	}

	if !userFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "User dengan ID tersebut tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User promoted to candidate", "user": updatedUser})
}
