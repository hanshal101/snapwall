package clickhouse

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var CHClient driver.Conn

func InitClickhouse(ctx context.Context) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{os.Getenv("CLICKHOUSE_ADDR")},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USERNAME"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
		Debug:           true,
		MaxOpenConns:    50,
		MaxIdleConns:    50,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		log.Fatalf("Error in starting Clickhouse Client: %v\n", err)
		return
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatalf("Error in Pinging the Clickhouse Client: %v\n", err)
		return
	}

	CHClient = conn
}
