package test

import (
	"bytes"
	"dating-app/internal/handlers"
	"dating-app/internal/middleware"
	"dating-app/internal/models"
	"dating-app/internal/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestViewProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.GET("/profiles", handlers.ViewProfiles)

	token, _ := utils.GenerateToken(1)
	req, _ := http.NewRequest(http.MethodGet, "/profiles", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSwipeProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/swipe", handlers.SwipeProfile)

	swipe := models.Swipe{
		ProfileID: 1,
		SwipeType: "like",
	}
	body, _ := json.Marshal(swipe)

	token, _ := utils.GenerateToken(1)
	req, _ := http.NewRequest(http.MethodPost, "/swipe", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Swipe recorded successfully")
}

func TestPurchasePremium(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/premium", handlers.PurchasePremium)

	premium := models.PremiumPackage{
		PackageType: "no_swipe_quota",
	}
	body, _ := json.Marshal(premium)

	token, _ := utils.GenerateToken(1)
	req, _ := http.NewRequest(http.MethodPost, "/premium", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Premium package purchased successfully")
}
