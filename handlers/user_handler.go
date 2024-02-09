package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Users retrieved successfully", "users": users})
}

func register(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" || newUser.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	// Check for duplicate username or email
	if isDuplicate(newUser.Name, newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email is already in use"})
		return
	}

	// Generate ID (you may use a library to generate unique IDs)
	newUser.ID = len(users) + 1
	if newUser.Role == "" {
		newUser.Role = "user"
	}

	users = append(users, newUser)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": newUser})
}

func login(c *gin.Context) {
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginInfo.Email == "" || loginInfo.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	user := findUserByEmail(loginInfo.Email)
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verifikasi kata sandi (gunakan library hashing untuk implementasi yang lebih aman)
	if user.Password != loginInfo.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Jika login berhasil, buat token JWT
	token, err := generateToken(user.ID)
	if err != nil {
		// Handle error
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	for i, u := range users {
		if u.ID == user.ID {
			users[i].Token = token
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func editProfile(c *gin.Context) {
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

	// Ambil data pengguna yang akan diedit
	var updatedUser User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi field yang tidak boleh kosong
	if updatedUser.Name == "" || updatedUser.Email == "" || updatedUser.Password == "" || updatedUser.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	for i, u := range users {
		if u.ID == userID {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password
			users[i].PhoneNumber = updatedUser.PhoneNumber
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "user": updatedUser})
}

func deleteUser(c *gin.Context) {
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
	if getUserRoleFromToken(c) != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden. Only admin can delete users"})
		return
	}

	// Ambil ID pengguna yang akan dihapus dari parameter URL
	ID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Temukan dan hapus pengguna berdasarkan ID
	index, existingUser := findUserByID(ID)
	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hapus pengguna dari slice users
	users = append(users[:index], users[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
