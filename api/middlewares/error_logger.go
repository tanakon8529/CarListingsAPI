package middlewares

import (
	"context"
	"daveslist-emdpcv/api/services"
	"log"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// ErrorLogger logs errors to Redis
func ErrorLogger() gin.HandlerFunc {
	redisClient := services.GetRedisClient()

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				err := redisClient.LPush(ctx, "errors", e).Err()
				if err != nil {
					log.Printf("Failed to log error to Redis: %v", err)
				}
			}
		}
	}
}
