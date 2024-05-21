package handlers

import (
	"dating-app/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ViewProfiles(c *gin.Context) {
	var profiles []models.Profile
	userID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 32)

	// Fetch 10 profiles that the user hasn't swiped today
	// (Assume we have a Swipe model to track swipes)

	// Example logic (you should implement the actual logic):
	models.DB.Where("user_id != ?", userID).Limit(10).Find(&profiles)

	c.JSON(http.StatusOK, profiles)
}

func SwipeProfile(c *gin.Context) {
	var swipe models.Swipe
	if err := c.ShouldBindJSON(&swipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swipe.UserID = uint(c.MustGet("user_id").(float64))
	swipe.CreatedAt = time.Now()

	if err := models.DB.Create(&swipe).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Swipe recorded successfully"})
}

func PurchasePremium(c *gin.Context) {
	var premium models.PremiumPackage
	if err := c.ShouldBindJSON(&premium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	premium.UserID = uint(c.MustGet("user_id").(float64))
	premium.ExpiryDate = time.Now().AddDate(0, 1, 0) // 1 month premium

	if err := models.DB.Create(&premium).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Premium package purchased successfully"})
}

func CreateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile.UserID = uint(c.MustGet("user_id").(float64))

	if err := models.DB.Create(&profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile created successfully", "profile": profile})
}

func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile.UserID = uint(c.MustGet("user_id").(float64))

	if err := models.DB.Model(&models.Profile{}).Where("user_id = ?", profile.UserID).Updates(profile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "profile": profile})
}
