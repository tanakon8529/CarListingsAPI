package controllers

import (
	"daveslist-emdpcv/api/database"
	"daveslist-emdpcv/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateListing godoc
// @Summary Create a new listing
// @Description Create a new listing
// @Tags Listings
// @Accept json
// @Produce json
// @Param listing body models.Listing true "Listing"
// @Success 200 {object} models.Listing
// @Failure 400 {object} map[string]interface{}
// @Router /listings [post]
func CreateListing(c *gin.Context) {
	var listing models.Listing
	if err := c.ShouldBindJSON(&listing); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	database.DB.Create(&listing)
	c.JSON(http.StatusOK, listing)
}

// GetListing godoc
// @Summary Get a listing by ID
// @Description Get a listing by ID
// @Tags Listings
// @Produce json
// @Param id path int true "Listing ID"
// @Success 200 {object} models.Listing
// @Failure 404 {object} map[string]interface{}
// @Router /listings/{id} [get]
func GetListing(c *gin.Context) {
	var listing models.Listing
	if err := database.DB.First(&listing, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Listing not found"})
		return
	}
	c.JSON(http.StatusOK, listing)
}

// UpdateListing godoc
// @Summary Update a listing by ID
// @Description Update a listing by ID
// @Tags Listings
// @Accept json
// @Produce json
// @Param id path int true "Listing ID"
// @Param listing body models.Listing true "Listing"
// @Success 200 {object} models.Listing
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /listings/{id} [put]
func UpdateListing(c *gin.Context) {
	var listing models.Listing
	if err := database.DB.First(&listing, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Listing not found"})
		return
	}
	if err := c.ShouldBindJSON(&listing); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	database.DB.Save(&listing)
	c.JSON(http.StatusOK, listing)
}

// DeleteListing godoc
// @Summary Delete a listing by ID
// @Description Delete a listing by ID
// @Tags Listings
// @Produce json
// @Param id path int true "Listing ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /listings/{id} [delete]
func DeleteListing(c *gin.Context) {
	if err := database.DB.Delete(&models.Listing{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Listing not found"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Listing deleted"})
}

// ListListings godoc
// @Summary List all listings
// @Description List all listings
// @Tags Listings
// @Produce json
// @Success 200 {object} []models.Listing
// @Router /listings [get]
func ListListings(c *gin.Context) {
	var listings []models.Listing
	database.DB.Find(&listings)
	c.JSON(http.StatusOK, listings)
}
