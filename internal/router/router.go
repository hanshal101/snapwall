package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/internal/application"
	"github.com/hanshal101/snapwall/internal/checkout"
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
	r.GET("/:ipAddr", policies.GetPoliciesbyIPs)
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

func CheckoutRoutes(r *gin.RouterGroup) {
	r.GET("/ips", checkout.GetCheckoutIPs)
	r.GET("/ips/:source", checkout.GetChkDetailsbyIPs)
	r.GET("/ips/:source/ports", checkout.GetChkIPsPorts)
}

func ApplicationRoutes(r *gin.RouterGroup) {
	r.GET("", application.GetApplications)
	r.POST("", application.CreateApplication)
	// r.PUT("/ports")
	r.DELETE("/:applicationID", application.DeleteApplication)
	r.GET("/port/:portNumber", application.GetApplicationsbyPort)
}
