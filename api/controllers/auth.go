package controllers

import (
	"daveslist-emdpcv/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TokenEndpoint godoc
// @Summary Generate a new token
// @Description Generate a new token if client_id and secret_id are correct
// @Tags auth
// @Accept  json
// @Produce  json
// @Param client_id header string true "Client ID"
// @Param secret_id header string true "Secret ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth [post]
func TokenEndpoint(c *gin.Context) {
	clientID := c.GetHeader("client_id")
	secretID := c.GetHeader("secret_id")

	token, err := services.GenerateToken(clientID, secretID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"type": "Bearer", "token": token})
}
