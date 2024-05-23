package handlers

import (
	"dating-app/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUnswipedProfiles(c *gin.Context) {
	var profiles []models.Profile
	userID := uint(c.MustGet("user_id").(float64))

	// Get today's date
	today := time.Now().Truncate(24 * time.Hour)
	var swipedProfileIDs []uint
	models.DB.Model(&models.Swipe{}).Where("user_id = ? AND swiped_at >= ?", userID, today).Pluck("profile_id", &swipedProfileIDs)
	if len(swipedProfileIDs) > 0 {
		models.DB.Where("id IN ?", swipedProfileIDs).Find(&profiles)
	} else {
		models.DB.Where("user_id != ?", userID).Find(&profiles)
	}
	c.JSON(http.StatusOK, profiles)
}

func ViewLikedProfiles(c *gin.Context) {
	var profiles []models.Profile
	userID := uint(c.MustGet("user_id").(float64))
	// Get today's date
	today := time.Now().Truncate(24 * time.Hour)

	// Fetch profile IDs that the user swiped today
	var swipedProfileIDs []uint
	models.DB.Model(&models.Swipe{}).Where("user_id = ? AND swiped_at >= ?", userID, today).Pluck("profile_id", &swipedProfileIDs)
	// Fetch 10 profiles that the user hasn't swiped today
	models.DB.Where("user_id IN ?", swipedProfileIDs).Limit(10).Find(&profiles)

	c.JSON(http.StatusOK, profiles)
}

func SwipeProfile(c *gin.Context) {
	var swipe models.Swipe
	if err := c.ShouldBindJSON(&swipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := uint(c.MustGet("user_id").(float64))
	swipe.UserID = userID
	swipe.SwipedAt = time.Now()

	// Get today's date
	today := time.Now().Truncate(24 * time.Hour)

	var premiumPackage models.PremiumPackage
	hasNoSwipeQuota := false
	if err := models.DB.Where("user_id = ? and package_type = ? and expiry_date >= ?", swipe.UserID, "no_swipe_quota", today).First(&premiumPackage).Error; err == nil {
		hasNoSwipeQuota = true
	}

	// Check if the user has already swiped on this profile today
	var existingSwipe models.Swipe
	if err := models.DB.Where("user_id = ? AND profile_id = ? AND swiped_at >= ?", swipe.UserID, swipe.ProfileID, today).First(&existingSwipe).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have already swiped on this profile today"})
		return
	}

	// Count the number of swipes the user has made today
	var swipeCount int64
	models.DB.Model(&models.Swipe{}).Where("user_id = ? AND swiped_at >= ?", swipe.UserID, today).Count(&swipeCount)

	if !hasNoSwipeQuota && swipeCount >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have reached the daily swipe limit"})
		return
	}

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
