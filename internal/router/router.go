package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/internal/logs"
	"github.com/hanshal101/snapwall/internal/policies"
)

func PolicyRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// Implement Policy Routes
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
	r.GET("", logs.GetLogs)
	r.GET("/port/:portNumber", logs.GetLogsByPort)
	r.GET("/:ioType/ip/:ipAddress", logs.GetLogsByIP)
	r.GET("/intruder", logs.GetIntruderLogs)
}
