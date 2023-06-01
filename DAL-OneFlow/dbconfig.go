package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetPGConnection(config ConnectionConfig) (*pgxpool.Pool, error) {
	connStr :=
		fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", config.Host, config.Port, config.Database, config.User, config.Password)
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Connection established")
	return pool, nil
}
