package controllers

import (
	"daveslist-emdpcv/api/database"
	"daveslist-emdpcv/api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SendPrivateMessage godoc
// @Summary Send a private message
// @Description Send a private message
// @Tags PrivateMessages
// @Accept json
// @Produce json
// @Param message body models.PrivateMessage true "Private Message"
// @Success 200 {object} models.PrivateMessage
// @Failure 400 {object} map[string]interface{}
// @Router /messages [post]
func SendPrivateMessage(c *gin.Context) {
	var message models.PrivateMessage
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	message.CreatedAt = time.Now()
	database.DB.Create(&message)
	c.JSON(http.StatusOK, message)
}

// GetPrivateMessage godoc
// @Summary Get a private message by ID
// @Description Get a private message by ID
// @Tags PrivateMessages
// @Produce json
// @Param id path int true "Message ID"
// @Success 200 {object} models.PrivateMessage
// @Failure 404 {object} map[string]interface{}
// @Router /messages/{id} [get]
func GetPrivateMessage(c *gin.Context) {
	var message models.PrivateMessage
	if err := database.DB.First(&message, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Message not found"})
		return
	}
	c.JSON(http.StatusOK, message)
}

// DeletePrivateMessage godoc
// @Summary Delete a private message by ID
// @Description Delete a private message by ID
// @Tags PrivateMessages
// @Produce json
// @Param id path int true "Message ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /messages/{id} [delete]
func DeletePrivateMessage(c *gin.Context) {
	if err := database.DB.Delete(&models.PrivateMessage{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Message not found"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"message": "Message deleted"})
}
