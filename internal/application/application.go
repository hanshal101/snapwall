package application

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/psql"
	"github.com/hanshal101/snapwall/models"
)

type CreateApplicationRequest struct {
	Name        string   `json:"name"`
	Port        string   `json:"port"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func GetApplications(c *gin.Context) {
	var applications []models.Application
	if err := psql.DB.Preload("Tags").Find(&applications).Error; err != nil {
		log.Fatalf("Error in fetching applications: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in fetching applications"})
		return
	}

	c.JSON(http.StatusOK, applications)
}

func CreateApplication(c *gin.Context) {
	var request CreateApplicationRequest
	if err := c.BindJSON(&request); err != nil {
		log.Fatalf("Error in binding request: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in binding request"})
		return
	}
	var application = models.Application{
		Name:        request.Name,
		Port:        request.Port,
		Description: request.Description,
	}

	var tags []models.Tags
	for _, tagReq := range request.Tags {
		tags = append(tags, models.Tags{
			ApplicationID: application.ID,
			Tag:           tagReq,
		})
	}

	tx := psql.DB.Begin()
	if err := tx.Create(&application).Error; err != nil {
		defer tx.Rollback()
		log.Fatalf("Error in creating application: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in creating application"})
		return
	}
	for _, tag := range tags {
		if err := tx.Create(&tag).Error; err != nil {
			log.Fatalf("Error in creating tag: %v\n", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "error in creating tag"})
			return
		}
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"success": "Application created successfully"})
}

func DeleteApplication(c *gin.Context) {
	id := c.Param("applicationID")
	if err := psql.DB.Where("id = ?", id).Delete(&models.Application{}).Error; err != nil {
		log.Fatalf("Error in deleting application: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in deleting application"})
		return
	}
}

func GetApplicationsbyPort(c *gin.Context) {
	portNumber := c.Param("portNumber")
	var applications []models.Application
	if err := psql.DB.Where("port = ?", portNumber).Preload("Tags").Find(&applications).Error; err != nil {
		log.Fatalf("Error in fetching applications: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in fetching applications"})
		return
	}
	c.JSON(http.StatusOK, applications)
}
