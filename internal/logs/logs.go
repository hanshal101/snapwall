package logs

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/models"
)

// Convert time.Time to ClickHouse compatible string
func formatTimestamp(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// This function is to store LOGS in Clickhouse
func StoreLogs(ctx context.Context, data *models.Log) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS service_logs (
			time DateTime,
			type String,
			source String,
			destination String,
			port String,
			protocol String
		) ENGINE = MergeTree()
		ORDER BY (time, source, destination)
		PRIMARY KEY (time, source, destination)
		PARTITION BY toYYYYMMDD(time)
	`

	if err := clickhouse.CHClient.Exec(ctx, createTableQuery); err != nil {
		log.Fatalf("Error creating table: %v", err)
		return err
	}

	formattedTime := formatTimestamp(data.Time)

	batch, err := clickhouse.CHClient.PrepareBatch(ctx, `
		INSERT INTO service_logs (time, type, source, destination, port, protocol) VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Fatalf("Error preparing batch insert statement: %v", err)
		return err
	}

	if err := batch.Append(formattedTime, data.Type, data.Source, data.Destination, data.Port, data.Protocol); err != nil {
		log.Fatalf("Error appending data to batch: %v", err)
		return err
	}

	if err := batch.Send(); err != nil {
		log.Fatalf("Error sending batch data: %v", err)
		return err
	}
	log.Println("Data inserted successfully")

	return nil
}

// This function is to fetch LOGS from Clickhouse
func GetLogs(c *gin.Context) {
	query := `
		SELECT time, type, source, destination, port, protocol
		FROM service_logs
	`
	rows, err := clickhouse.CHClient.Query(context.TODO(), query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error executing query"})
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
		); err != nil {
			log.Fatalf("Error scanning row: %v", err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Error scanning row"})
			return
		}

		logs = append(logs, logEntry)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Error iterating over rows"})
		return
	}

	c.JSON(http.StatusOK, logs)
}

func GetLogsByPort(c *gin.Context) {
	port := c.Query("portNumber")

	query := `
        SELECT time, type, source, destination, port, protocol
        FROM service_logs
        WHERE port = ?
    `

	rows, err := clickhouse.CHClient.Query(context.TODO(), query, port)
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

func GetLogsByIP(c *gin.Context) {
	ioType := c.Query("ioType")
	ipAddress := c.Query("ipAddress")

	query := fmt.Sprintf(`
        SELECT time, type, source, destination, port, protocol
        FROM service_logs
        WHERE %s = ?
    `, ioType)

	rows, err := clickhouse.CHClient.Query(context.TODO(), query, ipAddress)
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
