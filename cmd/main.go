package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/database/migrate"
	"github.com/hanshal101/snapwall/database/psql"
	"github.com/hanshal101/snapwall/internal/router"
	"github.com/joho/godotenv"
)

var (
	ctx = context.Background()
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error in loading '.env': %v", err)
		return
	}
	psql.InitDB()
	clickhouse.InitClickhouse(ctx)
	migrate.MigrateModels(psql.DB)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// POLICY Routes
	policy := r.Group("/policies")
	router.PolicyRoutes(policy)

	// LOG Routes
	log := r.Group("/logs")
	router.LogRoutes(log)

	r.Run(os.Getenv("APP_ADDRESS"))
}
