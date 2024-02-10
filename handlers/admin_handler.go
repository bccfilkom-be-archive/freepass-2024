package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"freepass-2024/models"
	"freepass-2024/utils"
)

func SetElectionDates(c *gin.Context) {
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

	var newSetting models.Setting
	if err := c.ShouldBindJSON(&newSetting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Jika tidak ada pengaturan tanggal pemilihan, tambahkan satu
	if len(models.Settings) == 0 {
		newSetting.ID = 1
		models.Settings = append(models.Settings, newSetting)
		c.JSON(http.StatusOK, gin.H{"message": "Election dates set successfully", "newSetting": newSetting})
		return
	}

	// Jika ada pengaturan tanggal pemilihan, ganti dengan yang baru
	models.Settings[0] = newSetting
	c.JSON(http.StatusOK, gin.H{"message": "Election dates updated successfully", "newSetting": newSetting})
}
func PromoteToCandidate(c *gin.Context) {
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

	// Mendapatkan candidateID dari URL
	candidateID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	var updatedUser models.User
	userFound := false
	// Mencari pengguna dengan ID yang sesuai dan memperbarui perannya menjadi "candidate"
	for i, u := range models.Users {
		if u.ID == candidateID {
			models.Users[i].Role = "candidate"
			models.Users[i].Token = ""
			updatedUser = models.Users[i]
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
