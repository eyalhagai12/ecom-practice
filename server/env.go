package server

import (
	"context"
	"ecom/db"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	DB      *pgxpool.Pool
	Queries *db.Queries
}

func NewEnv(ctx context.Context) Env {
	cfg := LoadCondiguration(ctx)
	pool := initDBPool(ctx, cfg.DbConnectionString)
	return Env{
		DB:      pool,
		Queries: initQueries(pool),
	}
}

func initDBPool(ctx context.Context, connectionString string) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatal("failed to connect to databse and init pool: ", err.Error())
	}

	return pool
}

func initQueries(pool *pgxpool.Pool) *db.Queries {
	queries := db.New(pool)

	return queries
}
