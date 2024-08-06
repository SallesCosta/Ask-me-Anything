package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sallescosta/ama/internal/api"
	"github.com/sallescosta/ama/internal/store/pgstore"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
	envPath, err := filepath.Abs("../../.env")
	if err != nil {
		panic(err)
	}

	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	ctx := context.Background()
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("WSRS_DATABASE_USER"),
		os.Getenv("WSRS_DATABASE_PASSWORD"),
		os.Getenv("WSRS_DATABASE_HOST"),
		os.Getenv("WSRS_DATABASE_PORT"),
		os.Getenv("WSRS_DATABASE_NAME"),
	)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {

		fmt.Println("2 - Aqui---------..")
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		if err := http.ListenAndServe(":8080", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
