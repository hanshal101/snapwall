package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall-backend/internal/policies"
)

func PolicyRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.GET("", policies.GetPolicies)
	r.POST("", policies.CreatePolicies)
	r.PUT("/:policyID", policies.UpdatePolicies)
	r.DELETE("/:policyID", policies.DeletePolicy)
}

func LogRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// Implement log routes
	// r.GET("")
	// r.GET("/port/:portNumber")
	// r.GET("/ip/:ipAddress")
	r.GET("")
}
