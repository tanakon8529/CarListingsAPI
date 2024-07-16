package controllers

import (
	"daveslist-emdpcv/api/database"
	"daveslist-emdpcv/api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateReply godoc
// @Summary Create a new reply
// @Description Create a new reply
// @Tags Replies
// @Accept json
// @Produce json
// @Param listing_id path int true "Listing ID"
// @Param reply body models.Reply true "Reply"
// @Success 200 {object} models.Reply
// @Failure 400 {object} map[string]interface{}
// @Router /listings/{listing_id}/replies [post]
func CreateReply(c *gin.Context) {
	var reply models.Reply
	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	reply.CreatedAt = time.Now()

	// Convert listing_id from string to uint
	listingID, err := strconv.ParseUint(c.Param("listing_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid listing ID"})
		return
	}
	reply.ListingID = uint(listingID)

	database.DB.Create(&reply)
	c.JSON(http.StatusOK, reply)
}

// GetReply godoc
// @Summary Get a reply by ID
// @Description Get a reply by ID
// @Tags Replies
// @Produce json
// @Param id path int true "Reply ID"
// @Success 200 {object} models.Reply
// @Failure 404 {object} map[string]interface{}
// @Router /replies/{id} [get]
func GetReply(c *gin.Context) {
	var reply models.Reply
	if err := database.DB.First(&reply, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Reply not found"})
		return
	}
	c.JSON(http.StatusOK, reply)
}

// UpdateReply godoc
// @Summary Update a reply by ID
// @Description Update a reply by ID
// @Tags Replies
// @Accept json
// @Produce json
// @Param id path int true "Reply ID"
// @Param reply body models.Reply true "Reply"
// @Success 200 {object} models.Reply
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /replies/{id} [put]
func UpdateReply(c *gin.Context) {
	var reply models.Reply
	if err := database.DB.First(&reply, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Reply not found"})
		return
	}
	if err := c.ShouldBindJSON(&reply); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	database.DB.Save(&reply)
	c.JSON(http.StatusOK, reply)
}

// DeleteReply godoc
// @Summary Delete a reply by ID
// @Description Delete a reply by ID
// @Tags Replies
// @Produce json
// @Param id path int true "Reply ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /replies/{id} [delete]
func DeleteReply(c *gin.Context) {
	if err := database.DB.Delete(&models.Reply{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Reply not found"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Reply deleted"})
}
