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
	r.GET("/:applicationID", policies.GetPolicies)
	r.POST("/:applicationID/:type", policies.CreatePolicies)
	r.PUT("/:applicationID/:policyID", policies.UpdatePolicies)
	r.DELETE("/:applicationID/:policyID", policies.DeletePolicy)
	// r.GET("/:ipAddr", policies.GetPoliciesbyIPs)
}

func LogRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// Implement log routes
	r.GET("/:applicationID", logs.GetLogs)
	r.GET("/port/:portNumber", logs.GetLogsByPort)
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

	r.GET("/:applicationID/policies")
	r.POST("/:applicationID/policies/:type")
	r.DELETE("/:applicationID/policies/:type/:policyID")
}
