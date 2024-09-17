package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hanshal101/snapwall/database/clickhouse"
	"github.com/hanshal101/snapwall/database/migrate"
	"github.com/hanshal101/snapwall/database/psql"
	"github.com/hanshal101/snapwall/internal/router"
	"github.com/hanshal101/snapwall/internal/sysinfo"
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

	r.GET("/sysinfo", sysinfo.GetSystemInfo)
	r.GET("/node", sysinfo.ServeNodeInfo)

	r.GET("/test", func(ctx *gin.Context) {
		bd, err := http.Get("http://192.168.200.1:5465")
		if err != nil {
			log.Fatalf("error : %v\n", err.Error())
			return
		}
		ctx.JSON(http.StatusOK, bd.Body)
	})
	// POLICY Routes
	policy := r.Group("/policies")
	router.PolicyRoutes(policy)

	// LOG Routes
	log := r.Group("/logs")
	router.LogRoutes(log)

	// CHECKOUT Routes
	checkout := r.Group("/checkout")
	router.CheckoutRoutes(checkout)

	// APPLICATION Routes
	application := r.Group("/application")
	router.ApplicationRoutes(application)

	r.Run(os.Getenv("APP_ADDRESS"))
}
