package server

import (
	"context"
	"ecom/db"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	Queries *db.Queries
}

func NewEnv(ctx context.Context) Env {
	cfg := LoadCondiguration(ctx)
	return Env{
		Queries: connectToDatabase(ctx, cfg.DbConnectionString),
	}
}

func connectToDatabase(ctx context.Context, connectionString string) *db.Queries {
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		log.Fatal("failed to connect to databse and init pool: ", err.Error())
	}

	queries := db.New(pool)

	return queries
}
