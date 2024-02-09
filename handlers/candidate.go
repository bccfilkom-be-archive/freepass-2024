package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafli5131/freepass-2024/models"
	"github.com/rafli5131/freepass-2024/utils"
)

func ViewCandidateInfo(c *gin.Context) {
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
	var candidates []gin.H
	for _, u := range models.Users {
		if u.Role == "candidate" {
			candidate := gin.H{
				"id":   u.ID,
				"name": u.Name,
			}
			candidates = append(candidates, candidate)
		}
	}

	c.JSON(http.StatusOK, gin.H{"candidates": candidates})
}
