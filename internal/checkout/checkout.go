package checkout

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/models"
)

func GetCheckoutIPs(c *gin.Context) {
	query := `
		SELECT DISTINCT source
		FROM service_logs;
	`

	rows, err := clickhouse.CHClient.Query(context.TODO(), query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing query"})
		return
	}
	defer rows.Close()

	var sources []string
	for rows.Next() {
		var source string

		if err := rows.Scan(
			&source,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		sources = append(sources, source)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving logs"})
		return
	}

	c.JSON(http.StatusOK, sources)
}

func GetChkDetailsbyIPs(c *gin.Context) {
	source := c.Param("source")

	query := `
        SELECT time, type, source, destination, port, protocol, severity
        FROM service_logs
        WHERE source = ? OR destination = ?
    `

	rows, err := clickhouse.CHClient.Query(context.TODO(), query, source, source)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing query"})
		return
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var logEntry models.Log

		if err := rows.Scan(
			&logEntry.Time,
			&logEntry.Type,
			&logEntry.Source,
			&logEntry.Destination,
			&logEntry.Port,
			&logEntry.Protocol,
			&logEntry.Severity,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		logs = append(logs, logEntry)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func GetChkIPsPorts(c *gin.Context) {
	source := c.Param("source")

	query := `
		SELECT DISTINCT port
		FROM service_logs
		WHERE source = ?
	`

	rows, err := clickhouse.CHClient.Query(context.TODO(), query, source)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing query"})
		return
	}
	defer rows.Close()

	var ports []string
	for rows.Next() {
		var port string

		if err := rows.Scan(
			&port,
		); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		ports = append(ports, port)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving logs"})
		return
	}

	c.JSON(http.StatusOK, ports)
}
