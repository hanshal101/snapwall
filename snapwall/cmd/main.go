package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database"
	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/internal"
	"github.com/joho/godotenv"
)

var (
	ctx = context.TODO()
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error in loading '.env': %v\n", err)
		return
	}
	clickhouse.InitClickhouse(ctx)
	database.InitDB()
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	// r.GET("/store", func(c *gin.Context) {
	// 	log.Println("Storing Sample Logs in Clickhouse")
	// 	internal.StoreData(ctx)
	// 	return
	// })
	r.GET("/get", func(c *gin.Context) {
		log.Println("Getting Sample Logs in Clickhouse")
		logs, err := internal.GetData(ctx)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		for i := range logs {
			log.Printf("Stored: %v\n", logs[i])
		}
		return
	})
	r.Run(":6545")
}
