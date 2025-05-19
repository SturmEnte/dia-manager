package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectPostgres(uri string) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, uri)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    if err := pool.Ping(ctx); err != nil {
        log.Fatalf("Unable to ping database: %v\n", err)
    }

    log.Println("Connected to PostgreSQL")
    DB = pool
}
