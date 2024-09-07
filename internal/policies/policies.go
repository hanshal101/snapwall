package policies

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/psql"
	"github.com/hanshal101/snapwall/internal/enforcer"
	"github.com/hanshal101/snapwall/models"
)

type PolicyRequest struct {
	Name  string   `json:"name"`
	IPs   []string `json:"ips"`
	Ports []string `json:"ports"`
	Type  string   `json:"type"`
}

func GetPolicies(c *gin.Context) {
	var policies []models.Policy
	if err := psql.DB.Preload("IPs").Preload("Ports").Find(&policies).Error; err != nil {
		log.Printf("Error in fetching policies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in fetching policies"})
		return
	}
	c.JSON(http.StatusOK, policies)
}

func CreatePolicies(c *gin.Context) {
	var req PolicyRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("Error in binding policies: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in binding policies"})
		return
	}

	tx := psql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	policy := models.Policy{
		Name: req.Name,
		Type: req.Type,
	}

	if err := tx.Create(&policy).Error; err != nil {
		tx.Rollback()
		log.Printf("Error in creating policy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating policy"})
		return
	}

	var ips []models.IP
	for _, ip := range req.IPs {
		ip := models.IP{
			PolicyID: policy.ID,
			Address:  ip,
		}
		if err := tx.Create(&ip).Error; err != nil {
			tx.Rollback()
			log.Printf("Error in creating IPs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating IPs"})
			return
		}
		ips = append(ips, ip)
	}

	var ports []models.Port
	for _, port := range req.Ports {
		pt := models.Port{
			PolicyID: policy.ID,
			Number:   port,
		}
		if err := tx.Create(&pt).Error; err != nil {
			tx.Rollback()
			log.Printf("Error in creating Ports: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating Ports"})
			return
		}
		ports = append(ports, pt)
	}

	if err := enforcer.ReconcileEnforcer(context.TODO(), policy, ips, ports); err != nil {
		log.Fatalf("Error in enforcement: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating Ports"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": "Policy Created Successfully"})
}

func UpdatePolicies(c *gin.Context) {
	var policyReq PolicyRequest
	if err := c.BindJSON(&policyReq); err != nil {
		log.Printf("Error in binding policies: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in binding policies"})
		return
	}

	policyID := c.Param("policyID")

	tx := psql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	var policy models.Policy
	if err := tx.Preload("IPs").Preload("Ports").First(&policy, policyID).Error; err != nil {
		tx.Rollback()
		log.Printf("Error fetching policy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching policy"})
		return
	}

	policy.Name = policyReq.Name
	policy.Type = policyReq.Type

	if err := tx.Save(&policy).Error; err != nil {
		tx.Rollback()
		log.Printf("Error saving policy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving policy"})
		return
	}

	if err := tx.Where("policy_id = ?", policyID).Delete(&models.IP{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting IPs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting IPs"})
		return
	}

	if err := tx.Where("policy_id = ?", policyID).Delete(&models.Port{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error deleting Ports: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Ports"})
		return
	}

	for _, ip := range policyReq.IPs {
		if err := tx.Create(&models.IP{PolicyID: policy.ID, Address: ip}).Error; err != nil {
			tx.Rollback()
			log.Printf("Error in creating IPs: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating IPs"})
			return
		}
	}

	for _, port := range policyReq.Ports {
		if err := tx.Create(&models.Port{PolicyID: policy.ID, Number: port}).Error; err != nil {
			tx.Rollback()
			log.Printf("Error in creating Ports: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in creating Ports"})
			return
		}
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": "Policy Updated Successfully"})
}

func DeletePolicy(c *gin.Context) {
	id := c.Param("policyID")

	var policy models.Policy
	if err := psql.DB.Where("id = ?", id).Preload("IPs").Preload("Ports").Find(&policy).Error; err != nil {
		log.Printf("Error in fetching policies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in fetching policies"})
		return
	}

	tx := psql.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}()

	if err := tx.Where("id = ?", id).Delete(&models.Policy{}).Error; err != nil {
		tx.Rollback()
		log.Printf("Error in deleting policy: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting policy"})
		return
	}

	var ips []models.IP
	if err := tx.Where("policy_id = ?", id).Delete(&models.IP{}).Find(&ips).Error; err != nil {
		tx.Rollback()
		log.Printf("Error in deleting IPs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting IPs"})
		return
	}

	var ports []models.Port
	if err := tx.Where("policy_id = ?", id).Delete(&models.Port{}).Find(&ports).Error; err != nil {
		tx.Rollback()
		log.Printf("Error in deleting Ports: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting Ports"})
		return
	}

	if err := enforcer.DeleteRule(context.TODO(), policy, policy.IPs, policy.Ports); err != nil {
		log.Fatalf("Error in Deleting rule: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error in deleting tables"})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"success": "Policy Deleted Successfully"})
}
