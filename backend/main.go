package main

import (
	"context"
	"example-backend/internal/customer"
	"example-backend/internal/repository"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load("./backend/config.env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get DB connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(ctx)

	// Create repository instance
	repository := repository.New(conn)

	customerService := customer.NewCustomerService(repository)

	customerService.Create(ctx, "Test", "Person", "Street 3", "testman", []byte("password"))
}
