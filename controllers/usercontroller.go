package controllers

import (
	"net/http"
	"pr/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("ubrtewtvyilyrb785b4o7w5yo8")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
    expirationTime := time.Now().Add(5 * time.Minute)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func Register(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password sebelum disimpan
    if err := input.HashPassword(input.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
        return
    }

    models.DB.Create(&input)
    c.JSON(http.StatusOK, gin.H{"data": "Registration successful"})
}

func Login(c *gin.Context) {
    var input models.User
    var user models.User

    // Mengambil data JSON yang dikirim oleh client
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cek apakah username ada di database
    if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
        return
    }

    // Cek kecocokan password
    if err := user.CheckPassword(input.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

    // Generate token JWT setelah login berhasil
    token, err := GenerateJWT(user.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
