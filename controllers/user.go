package controllers

import (
    "github.com/AbhijeetPundkar/ecommerce-backend/models"
    "github.com/AbhijeetPundkar/ecommerce-backend/middlewares"
    "github.com/AbhijeetPundkar/ecommerce-backend/internals/db"
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Name = strings.TrimSpace(user.Name)
    user.Username = strings.TrimSpace(user.Username)
    user.Email = strings.TrimSpace(user.Email)

    if err := user.HashPassword(user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    token, err := middlewares.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var dbUser models.User
    if err := db.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    if !dbUser.CheckPassword(user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := middlewares.GenerateToken(dbUser.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func Protected(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!", "userID": userID})
}