package common

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DBPool *pgxpool.Pool

func Init() *pgxpool.Pool {
	// connect to PriceTracker database
	pool, err := pgxpool.Connect(context.Background(), "postgresql://postgres:@localhost:5432/PriceTracker")

	if err != nil {
		fmt.Println("DB connection err: ", err)
	}

	DBPool = pool

	return DBPool
}

func GetDBPool() *pgxpool.Pool {
	return DBPool
}
